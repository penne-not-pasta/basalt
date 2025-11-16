package cmd

import (
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Hello",
	Short: "Welcomes the user",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Hello world!")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}


