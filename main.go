/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"os"

	"github.com/the-Jinxist/cmd_genie/cmd"
	"github.com/the-Jinxist/cmd_genie/config"
)

func main() {

	println("apiKey: ", os.Getenv("GEMINI_API_KEY"), "environment: ", os.Getenv("ENVIRONMENT"), "-------------0")

	config.LoadConfigs(".")
	cmd.Execute()
}
