package auth

import (
	"log"

	"github.com/spf13/cobra"
)

var AuthCmd = &cobra.Command{
	Use:   "auth",
	Short: "Login with cloud providers",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Auth command executed")
		log.Print("Auth Command executed")
	},
}
