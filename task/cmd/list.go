package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "View all the uncompleted tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is a fake task")
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
