package main

import (
	"fmt"
	"odeer/internal/config"
	"odeer/internal/conversation"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Println("⚠️ Error loading configuration:", err)
		return
	}

	conversation.RunConversation(config)
}
