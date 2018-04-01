package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show all the available Tasks",
	Long:  `Show all the available Tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Here goes the unresolved tasklist: ")
	},
}