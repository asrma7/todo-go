/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strconv"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove <task_id>",
	Short: "Remove a todo task from the list by its id.",
	Long: `Remove a todo task from the list by its id. For example:
    todo remove 1`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide the task ID to remove.")
			return
		} else if len(args) > 1 {
			fmt.Println("Please provide only one task ID to remove.")
			return
		}

		todoId, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid task ID. Please provide a numeric task ID.")
			return
		}

		if todoId < 0 || todoId >= len(todos.Todos) {
			fmt.Printf("Task ID %d is out of range. Please provide a valid task ID.\n", todoId)
			return
		}

		// Remove the task from the list
		todos.Todos = append(todos.Todos[:todoId], todos.Todos[todoId+1:]...)

		// Update the configuration
		viper.Set("todos", todos.Todos)
		err = viper.WriteConfig()
		if err != nil {
			fmt.Println("Error writing to config file:", err)
			return
		}

		fmt.Printf("Task %d removed.\n", todoId)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
