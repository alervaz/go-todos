/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"todos/model"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new todo list",
	Long:  `This will create a new record in the database that requires a description`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Error("You need to specify todo")
			return
		}

		for _, todo := range args {
			if todo == "not completed" || todo == "completed" {
				log.Error(`Todo cannot be "completed" or "not completed"`)
			}

			tsx := model.DB.Create(&model.Todo{
				Todo:      todo,
				Completed: false,
			})

			if tsx.Error != nil {
				log.Error(tsx.Error)
				continue
			}

			log.Info(todo + " was succesfully added to the list")
		}

	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
