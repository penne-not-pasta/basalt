package cmd

import (
	"log"
	"oreonproject/basalt/cmd/auth"
	"oreonproject/basalt/utils"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bassalt",
	Short: "Root command",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Hello world!")
		log.Print("Root Command Executed")
	},
}

func Execute() {
	utils.LogSetup("logs/root.log")
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(auth.AuthCmd)
	auth.AuthCmd.AddCommand(auth.GoogleCmd)
}
