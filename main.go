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

	// Configure the model
	model := client.GenerativeModel("gemini-1.5-flash-8b")
	model.SetTemperature(1)
	model.SetTopK(40)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(8192)
	model.ResponseMIMEType = "text/plain"
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{genai.Text("Please provide a text that needs to be fixed. Ensure proper grammar, punctuation, and clarity.")},
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
