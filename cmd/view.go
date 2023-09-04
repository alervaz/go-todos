/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strconv"
	"todos/model"

	"github.com/charmbracelet/log"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View the todo list",
	Long:  `This will show a the records in the database in detail`,
	Run: func(cmd *cobra.Command, args []string) {
		var todos []model.Todo

		state, err := cmd.Flags().GetString("state")
		if err != nil {
			log.Error(err)
			return
		}

		id, _ := cmd.Flags().GetInt("id")
		if err != nil {
			log.Error(err)
			return
		}
		if state != "" && state != "not completed" && state != "completed" {
			log.Warn("Incorrect state default state selected")
		}

		if id != 0 {
			tsx := model.DB.Where("id = ?", id).Find(&todos)
			if tsx.Error != nil || len(todos) == 0 {
				log.Error("Todo does not exist")
				return
			}
		} else if state == "not completed" {
			tsx := model.DB.Where("completed = ?", 0).Find(&todos)
			if tsx.Error != nil {
				log.Error(tsx.Error)
				return
			}

		} else if state == "completed" {
			tsx := model.DB.Where("completed = ?", 1).Find(&todos)
			if tsx.Error != nil {
				log.Error(tsx.Error)
				return
			}

		} else {

			tsx := model.DB.Find(&todos)
			if tsx.Error != nil {
				log.Error(tsx.Error)
				return
			}

		}

		if len(todos) == 0 {
			log.Warn("There are not todos in the list")
			return
		}

		headerFmt := color.New(color.FgCyan, color.Underline).SprintfFunc()
		columnFmt := color.New(color.FgBlue).SprintfFunc()

		tbl := table.New("ID", "Todo", "Completed", "CreatedAt")
		tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

		for _, todo := range todos {
			todoText := chalk.Red.Color(todo.Todo)
			completed := chalk.Red.Color(strconv.FormatBool(todo.Completed))

			if todo.Completed {
				todoText = chalk.Green.Color(todo.Todo)
				completed = chalk.Green.Color(strconv.FormatBool(todo.Completed))
			}

			tbl.AddRow(todo.ID, todoText, completed, chalk.Blue.Color(todo.CreatedAt.String()))
		}

		tbl.Print()
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
	rootCmd.PersistentFlags().String("state", "", "Search by completion")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// viewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// viewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
