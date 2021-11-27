package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
	billy "github.com/wimgoeman/billy/core"
)

var newInvoiceCmd = &cobra.Command{
	Use: "invoice",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("new invoice called")
	},
}

var sendInvoiceCmd = &cobra.Command{
	Use:  "invoice",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		result, err := billy.SendInvoice("default", args[0])
		cobra.CheckErr(err)
		fmt.Println("Generated invoice at: " + result)
		err = exec.Command("chromium", result).Start()
		cobra.CheckErr(err)
	},
}

func init() {
	newCmd.AddCommand(newInvoiceCmd)
	sendCmd.AddCommand(sendInvoiceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// invoiceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// invoiceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
