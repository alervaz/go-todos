/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/charmbracelet/log"
	"os"
	"todos/cmd"
	"todos/model"
)

func init() {
	if err := model.InitializeDB(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}

func main() {

	cmd.Execute()
}
