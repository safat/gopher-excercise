package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: `Add a new Task to the existing list`,
	Long:  `Add a new Task to the existing list`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Adding a new task")
	},
}
