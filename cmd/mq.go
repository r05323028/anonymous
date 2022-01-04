package cmd

import "github.com/spf13/cobra"

var mqCmd = &cobra.Command{
	Use:   "mq",
	Short: "Message Queue Command Line Tools",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(mqCmd)
}
