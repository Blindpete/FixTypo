# FixTypo CLI

A simple command-line chat application powered by Google's Gemini AI model (Gemini 1.5 Flash 8B) to help improve text by focusing on grammar and clarity directly through your terminal.


## To-Do
- Add: Multi-line text support
- Add: Copy to the clipboard

## Prerequisites

- Go 1.21 or higher
- A Google AI Studio API key (Gemini)

## Installation

You can install the application directly using Go:

```bash
go install github.com/Blindpete/FixTypo@v0.2.1
```

Or clone and build manually:

```bash
git clone https://github.com/Blindpete/FixTypo.git
cd FixTypo
go build
```

## Configuration

Before running the application, you need to set up your Gemini API key as an environment variable:

### On Linux / macOS

```bash
# Get the API key from https://aistudio.google.com/apikey
export GEMINI_API_KEY='your-api-key-here'
```

### On Windows (PowerShell)

```powershell
# Get the API key from https://aistudio.google.com/apikey
$env:GEMINI_API_KEY = "your-api-key-here"
```

## Usage

After installation, simply run:

```bash
fixtypo
```

Or if built manually:

```bash
./fixtypo
```

### Commands

- Type your message and press Enter to send it to the Gemini AI for text improvement.
- Type `exit` to end the conversation.

## Model Configuration

The application uses the following default settings for the Gemini model:

- Model: `gemini-2.0-flash-lite`
- Temperature: 1.0
- Top-K: 40
- Top-P: 0.95
- Max Output Tokens: 8192
- Response Format: Plain text
- System Instruction: Focus on grammar, punctuation, and clarity improvements

## Example Interaction

```
Welcome to the FixTypo! Type 'exit' to end.
--- Thursday, November 7, 2024 at 22:08
🤠: lte's go
✨: Let's go!
```

## Dependencies

- [github.com/google/generative-ai-go/genai](https://github.com/google/generative-ai-go/genai)
- [google.golang.org/api](https://pkg.go.dev/google.golang.org/api)

## Error Handling

The application includes error handling for:
- Missing API key
- Client initialization failures
- Message sending errors
- Input scanning errors




