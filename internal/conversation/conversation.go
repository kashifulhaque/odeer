package conversation

import (
	"fmt"
	"log"
	"strings"

	"odeer/internal/api"
	"odeer/internal/config"
	"odeer/internal/models"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type (
	errMsg error
)

var messages = []models.Message{
	{
		Role:    "system",
		Content: "You are a helpful assistant, called EnderGPT",
	},
}

var globalConfig = config.Config{}

type apiResponseMsg string

func (m model) sendRequest() tea.Msg {
	assistantResponse, err := api.SendRequest(&globalConfig, messages)
	if err != nil {
		return errMsg(err)
	}
	messages = append(messages, models.Message{Role: "assistant", Content: assistantResponse})
	return apiResponseMsg(assistantResponse)
}

type model struct {
	viewport    viewport.Model
	messages    []string
	textarea    textarea.Model
	senderStyle lipgloss.Style
	err         error
	spinner     spinner.Model
	loading     bool
}

func initialModel() model {
	ta := textarea.New()
	ta.Placeholder = "Chat with EnderGPT ..."
	ta.Focus()

	ta.Prompt = "┃ "
	ta.CharLimit = 640

	ta.SetHeight(2)
	ta.SetWidth(120)

	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()

	ta.ShowLineNumbers = false

	vp := viewport.New(500, 50)
	vp.SetContent(`🦌 odeer`)

	ta.KeyMap.InsertNewline.SetEnabled(false)

	sp := spinner.New()
	sp.Spinner = spinner.Points
	sp.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return model{
		textarea:    ta,
		messages:    []string{},
		viewport:    vp,
		senderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("5")),
		err:         nil,
		spinner:     sp,
		loading:     false,
	}
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit

		case tea.KeyEnter:
			if !m.loading {
				userMessage := m.textarea.Value()
				m.messages = append(m.messages, m.senderStyle.Render("👤 ")+userMessage)
				m.viewport.SetContent(strings.Join(m.messages, "\n"))
				m.textarea.Reset()
				m.viewport.GotoBottom()

				messages = append(messages, models.Message{Role: "user", Content: userMessage})

				m.loading = true
				cmds = append(cmds, m.spinner.Tick)
				cmds = append(cmds, m.sendRequest)
			}

		default:
			if !m.loading {
				var cmd tea.Cmd
				m.textarea, cmd = m.textarea.Update(msg)
				cmds = append(cmds, cmd)
			}
		}

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		if m.loading {
			cmds = append(cmds, cmd)
		}

	case apiResponseMsg:
		m.loading = false
		assistantResponse := string(msg)
		messages = append(messages, models.Message{Role: "assistant", Content: assistantResponse})
		m.messages = append(m.messages, m.senderStyle.Render("✨ ")+assistantResponse)
		m.viewport.SetContent(strings.Join(m.messages, "\n"))
		m.viewport.GotoBottom()

	case errMsg:
		m.loading = false
		m.err = msg
		m.messages = append(m.messages, m.senderStyle.Render("⚠️  Error: ")+m.err.Error())
		m.viewport.SetContent(strings.Join(m.messages, "\n"))
		m.viewport.GotoBottom()

	default:
		var cmd tea.Cmd
		m.viewport, cmd = m.viewport.Update(msg)
		cmds = append(cmds, cmd)
	}

	if !m.loading {
		var vpCmd tea.Cmd
		m.viewport, vpCmd = m.viewport.Update(msg)
		cmds = append(cmds, vpCmd)
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	var content string
	if m.loading {
		content = fmt.Sprintf("%s Hold on, I am thinking ...", m.spinner.View())
	} else {
		content = m.viewport.View()
	}

	return fmt.Sprintf(
		"%s\n\n%s",
		content,
		m.textarea.View(),
	) + "\n\n"
}

func Run(cfg *config.Config) {
	p := tea.NewProgram(initialModel())
	globalConfig = *cfg

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
