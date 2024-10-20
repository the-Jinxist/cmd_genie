package ui

// A simple program demonstrating the text input component from the Bubbles
// component library.

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/the-Jinxist/cmd_genie/config"
	"github.com/the-Jinxist/cmd_genie/internal/chat_client"
)

type (
	errMsg error
)

type model struct {
	textInput textinput.Model
	err       error
}

func InitialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Command to copy text"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return model{
		textInput: ti,
		err:       nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			fmt.Println(strings.ToUpper(m.textInput.Value()))

			client := chat_client.NewClient(config.GetGeminiAPIKey())
			fmt.Println(fmt.Sprintf("\n \n Your API KEY: %s", client.APIKey))

			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"You need a command to do what?\n\n%s\n\n%s",
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}
