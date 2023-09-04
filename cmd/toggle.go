/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"todos/model"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

// toggleCmd represents the toggle command
var toggleCmd = &cobra.Command{
	Use:   "toggle",
	Short: "Toggle between completed and not completed",
	Long: `Toggle between completed and not completed use as many args as you want.
  
  to toggle pass the todo ID`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Error("Please specify todo ID")
			return
		}

		for _, id := range args {
			var todo model.Todo
			tsx := model.DB.Where("id = ?", id).First(&todo)
			if tsx.Error != nil {
				log.Error("Todo " + id + " does not exist")
				return
			}

			tsx = model.DB.Save(&model.Todo{
				Todo:      todo.Todo,
				ID:        todo.ID,
				Completed: !todo.Completed,
			})

			if tsx.Error != nil {
				log.Error(id + ": " + tsx.Error.Error())
			}
		}

		log.Info("Succesfully updated todos!")
	},
}

func init() {
	rootCmd.AddCommand(toggleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// toggleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// toggleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
