/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Long:  `WAL`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add a new task to the todo-list")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
