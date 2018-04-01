package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a light weight tool to manage your taks",
	Long:  `Task provides very simple functionality to add/remove/complete tasks. Internally it uses cobra cli manager.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
