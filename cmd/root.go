/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/the-Jinxist/cmd_genie/ui"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "cmd_genie",
	Aliases: []string{"genie"},
	Short:   "This application helps to find CLI commands that user's need",
	Long:    `It's simple you ask me for a command you can't seem to remember and I'll give you the closest suggestion`,
	Run: func(cmd *cobra.Command, args []string) {
		style := lipgloss.NewStyle().Padding(1, 2).Foreground(lipgloss.Color("#ffffff")).Border(lipgloss.DoubleBorder(), true)
		fmt.Println(style.Render("Welcome to your Command-Line Genie!"))

		p := tea.NewProgram(ui.InitialModel())
		if _, err := p.Run(); err != nil {
			log.Fatal(err)
		}
	},
}

var shortHandSuggestion = &cobra.Command{
	Use:     "get",
	Short:   "Short-hand to get cli command suggestions without opening the text input",
	Example: "cmd_genie get command to sign and create a git commit",
	Args:    cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		sentence := strings.Join(args, " ")

		p := tea.NewProgram(ui.InitalLoadingModel(sentence))
		if _, err := p.Run(); err != nil {
			log.Fatal(err)
		}

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(shortHandSuggestion)
}
