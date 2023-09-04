/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"todos/model"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes the completed todos",
	Long:  `Deletes the completed todos and can pass a flag with a specific todo`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetInt("id")
		if err != nil {
			log.Error(err)
			return
		}

		if id != 0 {
			tsx := model.DB.Where("id = ?", id).Unscoped().Delete(&model.Todo{})
			if tsx.Error != nil {
				log.Error("Todo does not exist")
				return
			}
			log.Info("Succesfully removed todo: ", id)
			return
		}

		tsx := model.DB.Where("completed = ?", "1").Unscoped().Delete(&model.Todo{})

		if tsx.Error != nil {
			log.Error(tsx.Error)
			return
		}

		log.Info("Succesfully removed todos")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	rootCmd.PersistentFlags().Int("id", 0, "To specify which todo will be deleted with the id")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
