package ui

// A simple program demonstrating the text input component from the Bubbles
// component library.

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/the-Jinxist/cmd_genie/config"
	"github.com/the-Jinxist/cmd_genie/internal/chat_client"
)

type (
	errMsg error
)

type model struct {
	textInput        textinput.Model
	err              error
	loading, success bool
	spinner          spinner.Model
}

func InitialModel() model {

	ti := textinput.New()
	ti.Placeholder = "Command to copy text"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return model{
		textInput: ti,
		spinner:   s,
		err:       nil,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(textinput.Blink, m.spinner.Tick)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:

			// fmt.Println(fmt.Sprintf("\n \n Your API KEY: %s", client.APIKey))
			prompt := m.textInput.Value()
			m.loading = true
			m.success = false

			return m, makeAPICall(prompt)
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		m.loading = false
		return m, nil
	case successMsg:
		m.success = true
		return m, nil
	}

	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)

	cmds = append(cmds, cmd)
	m.spinner, cmd = m.spinner.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	initStr := fmt.Sprintf(
		"You need a command to do what?\n\n%s\n\n%s",
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"

	if m.loading {
		initStr += fmt.Sprintf("\nChecking the matrix to find your command.. %s", m.spinner.View())
	}

	if m.success {
		initStr += "\nSuccessfully got command"
	}

	return initStr
}

type successMsg int

func makeAPICall(prompt string) tea.Cmd {

	return func() tea.Msg {
		client := chat_client.NewClient(config.GetGeminiAPIKey())
		resp, err := client.ChatCompletion.GetChatCompletion(prompt)
		if resp != nil {
			log.Default().Printf("Logger: %v\n", resp)
			return successMsg(len(resp.Candidates))
		}

		return errMsg(err)
	}

}
