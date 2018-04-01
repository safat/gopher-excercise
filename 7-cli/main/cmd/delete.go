package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: `Delete a Task from the existing list`,
	Long:  `Delete a Task from the existing list`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Deleting a Task")
	},
}
