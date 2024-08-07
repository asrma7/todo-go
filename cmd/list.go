/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"text/tabwriter"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks in the todo list.",
	Long: `List all tasks in the todo list. For example:
    todo list
    todo list --all
    todo list --completed
    todo list --all --completed`,

	Run: func(cmd *cobra.Command, args []string) {
		minwidth := 0
		tabwidth := 8
		padding := 1
		padchar := byte('\t')

		w := tabwriter.NewWriter(os.Stdout, minwidth, tabwidth, padding, padchar, tabwriter.AlignRight)

		if len(args) > 0 {
			fmt.Println("No arguments required for this command.")
			return
		}
		all, _ := cmd.Flags().GetBool("all")
		completed, _ := cmd.Flags().GetBool("completed")

		if all && completed {
			fmt.Println("Listing all tasks including the completed ones.")
			fmt.Fprintln(w, "ID\tTask\tDescription\tDeadline\tCompleted")
			fmt.Fprintln(w, "--\t----\t-----------\t--------\t---------")
			for id, todo := range todos.Todos {
				fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%t\n", id, todo.Name, todo.Description, todo.Deadline, todo.Completed)
			}
		} else if all {
			fmt.Println("Listing all tasks with deadlines and descriptions.")
			fmt.Fprintln(w, "ID\tTask\tDescription\tDeadline")
			fmt.Fprintln(w, "--\t----\t-----------\t--------")
			for id, todo := range todos.Todos {
				fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", id, todo.Name, todo.Description, todo.Deadline)
			}
		} else if completed {
			fmt.Println("Listing completed tasks.")
			fmt.Fprintln(w, "ID\tTask\tCompleted")
			fmt.Fprintln(w, "--\t----\t---------")
			for id, todo := range todos.Todos {
				if todo.Completed {
					fmt.Fprintf(w, "%d\t%s\t%t\n", id, todo.Name, todo.Completed)
				}
			}
		} else {
			fmt.Println("Listing pending tasks.")
			fmt.Fprintln(w, "ID\tTask")
			fmt.Fprintln(w, "--\t----")
			for id, todo := range todos.Todos {
				if !todo.Completed {
					fmt.Fprintf(w, "%d\t%s\n", id, todo.Name)
				}
			}
		}
		w.Flush()

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	listCmd.Flags().BoolP("all", "a", false, "List all tasks along with deadlines and descriptions")
	listCmd.Flags().BoolP("completed", "c", false, "List all tasks including the completed ones")
}
