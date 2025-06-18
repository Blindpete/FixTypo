package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()

	// Get API key from environment
	apiKey, ok := os.LookupEnv("GEMINI_API_KEY")
	if !ok {
		log.Fatalln("Environment variable GEMINI_API_KEY not set")
	}

	// Initialize Gemini client
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}
	defer client.Close()
	// Text correction prompt
	// var textCorrectionPrompt string = `Input text for correction. Apply British English conventions for grammar, spelling, and punctuation. Ensure the output is clear, concise, and adheres to the established style guide.`
	// var textCorrectionPromptPreserveEmoji string = `Input text for correction. Apply British English conventions for grammar, spelling, and punctuation. Retain all original emojis in their positions. Ensure the final text is clear, concise, and adheres to the established style guide.`
	// var textCorrectionPromptModernSemicolons string = `Input text for correction. Apply British English conventions for grammar, spelling, and punctuation. For a modern casual style, replace semicolons with alternatives where appropriate for clarity and flow. Retain all original emojis in their positions. Ensure the final text is clear, concise, and adheres to the established style guide.`
	var textCorrectionPromptModernSemicolons string = `Input text for correction. Apply British English conventions for grammar, spelling, and punctuation. For a modern casual style, replace semicolons with alternatives where appropriate for clarity and flow. Add appropriate emojis where they naturally enhance the message. Ensure the final text is clear, concise, and adheres to the established style guide.`
	// 	// Style Guide
	// 	var correctionStyleGuide string = `* Apply British English spelling, grammar, and punctuation.
	// * Ensure the text is clear and concise.
	// * Use active voice where it improves clarity and is grammatically appropriate.
	// * Prioritise accuracy and natural phrasing according to British English conventions.`

	// Configure the model
	model := client.GenerativeModel("gemini-2.5-flash-lite-preview-06-17")
	model.SetTemperature(0.7)
	model.SetTopK(40)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(8192)
	model.ResponseMIMEType = "text/plain"
	model.SystemInstruction = &genai.Content{
		// Parts: []genai.Part{genai.Text("Please provide a text that needs to be fixed. Ensure proper grammar, punctuation, and clarity.")},
		// Parts: []genai.Part{genai.Text(textCorrectionPrompt + "\n" + correctionStyleGuide)},
		Parts: []genai.Part{genai.Text(textCorrectionPromptModernSemicolons)},
	}

	// Start chat session
	session := model.StartChat()
	session.History = []*genai.Content{}

	// Create scanner for user input
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the FixTypo! Type 'exit' to end.")

	// Main chat loop
	for {
		currentTime := time.Now().Format("Monday, January 2, 2006 at 15:04")
		fmt.Println("--- " + currentTime)
		// Print prompt and get user input
		fmt.Print("🤠: ")
		if !scanner.Scan() {
			break
		}

		userInput := scanner.Text()

		// Check for exit command
		if strings.ToLower(strings.TrimSpace(userInput)) == "exit" {
			fmt.Println("\nGoodbye!")
			break
		}

		// Send message to Gemini
		resp, err := session.SendMessage(ctx, genai.Text(userInput))
		if err != nil {
			fmt.Printf("Error sending message: %v\n", err)
			continue
		}

		// Print AI response
		fmt.Print("✨: ")
		for _, part := range resp.Candidates[0].Content.Parts {
			fmt.Printf("%v", part)
		}
		fmt.Println()
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading input: %v", err)
	}
}
