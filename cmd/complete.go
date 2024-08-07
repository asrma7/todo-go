/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete <task_id>",
	Short: "Mark a todo task as completed.",
	Long: `Mark a todo task as completed. For example:
    todo complete 1`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide the task id to complete.")
			return
		} else if len(args) > 1 {
			fmt.Println("Please provide only one task id to complete.")
			return
		}
		todoId, err := strconv.Atoi(args[0]) // Get the task id from the command line arguments
		if err != nil {
			fmt.Println("Invalid task ID. Please provide a numeric task ID.")
			return
		}

		if todoId < 0 || todoId >= len(todos.Todos) {
			fmt.Printf("Task ID %d is out of range. Please provide a valid task ID.\n", todoId)
			return
		}
		todos.Todos[todoId].Completed = true

		viper.Set("todos", todos.Todos)
		err = viper.WriteConfig()

		if err != nil {
			fmt.Println("Error writing to config file:", err)
			return
		}
		fmt.Printf("Task %d marked as completed.\n", todoId)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
