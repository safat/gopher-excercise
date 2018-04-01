package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(doCmd)
}

var doCmd = &cobra.Command{
	Use:   "do",
	Short: `Mark a Task as completed from the existing list`,
	Long:  `Mark a Task as completed from the existing list`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Marking as completed")
	},
}
