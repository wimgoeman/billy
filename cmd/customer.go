package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var customerCmd = &cobra.Command{
	Use: "customer",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This will create a new customer at some point")
	},
}

func init() {
	newCmd.AddCommand(customerCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// customerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// customerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
