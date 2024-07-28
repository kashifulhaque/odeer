package conversation

import (
	"bufio"
	"fmt"
	"odeer/internal/api"
	"odeer/internal/config"
	"odeer/internal/models"
	"os"
	"strings"
)

func Run(config *config.Config) {
	messages := []models.Message{
		{
			Role:    "system",
			Content: "You are a helpful assistant, called EnderGPT",
		},
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("🦌 odeer (To exit, either type /bye or /exit).\nPowered by LLaMa 3 🦙\n\n")
	for {
		fmt.Print("❔ ")
		userInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("⚠️  Error reading input: %v\n", err)
			return
		}
		userInput = strings.ToLower(strings.TrimSpace(userInput))

		if userInput == "/bye" || userInput == "/exit" {
			fmt.Println("👋 Goodbye!")
			return
		}

		if userInput == "" {
			fmt.Println("⚠️  Hmm, seems like an empty string")
			continue
		}

		messages = append(messages, models.Message{Role: "user", Content: userInput})

		assistantResponse, err := api.SendRequest(config, messages)
		if err != nil {
			fmt.Printf("⚠️  Error getting response: %v\n", err)
			continue
		}

		messages = append(messages, models.Message{Role: "assistant", Content: assistantResponse})
		fmt.Println() // Add a newline for better readability
	}
}
