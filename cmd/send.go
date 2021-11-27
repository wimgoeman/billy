package cmd

import (
	"github.com/spf13/cobra"
)

var sendType string

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send a document",
}

func init() {
	sendCmd.Flags().StringVarP(&sendType, "type", "t", "html", "how you want to send the document")
	rootCmd.AddCommand(sendCmd)
}
