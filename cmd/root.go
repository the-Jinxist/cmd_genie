/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/the-Jinxist/cmd_genie/ui"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cmd_genie",
	Short: "This application helps to find CLI commands that user's need",
	Long:  `It's simple you ask me for a command you can't seem to remember and I'll give you the closest suggestion`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to your Command-Line Genie!")

		p := tea.NewProgram(ui.InitialModel())
		if _, err := p.Run(); err != nil {
			log.Fatal(err)
		}
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
	// rootCmd.AddCommand(suggestCommandCMD)
}
