/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add \"my todo task\"",
	Short: "Add a new todo task to the list.",
	Long: `Add a new todo task to the list. For example:
    todo add "Call mom" -d "I need to call mom for her birthday"
    todo add "Buy groceries" -d "I need to buy groceries before the weekend" -n "2024-12-31"`,

	Run: func(cmd *cobra.Command, args []string) {

		description, _ := cmd.Flags().GetString("description")
		deadline, _ := cmd.Flags().GetString("deadline")

		if len(args) == 0 {
			fmt.Println("Please provide the task name to add.")
			return
		} else if len(args) > 1 {
			fmt.Println("Please provide only one task name to add.")
			return
		} else if len(description) == 0 {
			fmt.Println("Please provide the description of the task.")
			return
		}

		todo := Todo{
			// Id:          len(todos.Todos),
			Name:        args[0],
			Description: description,
			Deadline:    deadline,
			Completed:   false,
		}

		todos.Todos = append(todos.Todos, todo)

		viper.Set("todos", todos.Todos)
		err := viper.WriteConfig()

		if err != nil {
			fmt.Println("Error writing to config file:", err)
			return
		}
		fmt.Printf("Task %d added.\n", len(todos.Todos)-1)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().StringP("description", "d", "", "Description of the todo task")
	addCmd.Flags().StringP("deadline", "n", "", "Deadline of the todo task")
	addCmd.MarkFlagRequired("description")
}
