/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/the-Jinxist/cmd_genie/cmd"
	"github.com/the-Jinxist/cmd_genie/config"
)

func main() {

	config.LoadConfigs(".")
	cmd.Execute()
}
