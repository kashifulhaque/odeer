package conversation

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"odeer/internal/api"
	"odeer/internal/config"
	"odeer/internal/models"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	viewport    viewport.Model
	messages    []models.Message
	textarea    textarea.Model
	senderStyle lipgloss.Style
	err         error
}

func initialModel() model {
	ta := textarea.New()
	ta.Placeholder = "Say something ..."
	ta.Focus()

	ta.Prompt = "┃ "
	ta.CharLimit = 500

	ta.SetWidth(30)
	ta.SetHeight(3)

	// Remove cursor line styling
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()

	ta.ShowLineNumbers = false

	vp := viewport.New(30, 5)
	vp.SetContent(`🦌 odeer
	Type a message and press Enter to send.`)
	ta.KeyMap.InsertNewline.SetEnabled(false)

	return model{
		textarea:    ta,
		messages:    []models.Message,
		viewport:    vp,
		senderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("5")),
		err:         nil,
	}
}

func Run(config *config.Config) {
	p := tea.NewProgram(initialModel())
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
