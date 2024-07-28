package main

import (
	"fmt"
	"odeer/internal/config"
	"odeer/internal/conversation"
	"odeer/internal/server"
	"os"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("⚠️  Error loading configuration:", err)
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: odeer [start|talk]")
		return
	}

	switch os.Args[1] {
	case "start":
		server.Start(cfg)
	case "talk":
		conversation.Run(cfg)
	default:
		fmt.Println("⚠️  Invalid command. Use 'start' or 'talk'")
	}
}
