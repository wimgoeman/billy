package cmd

import (
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a resource",
}

func init() {
	rootCmd.AddCommand(newCmd)
}
