package auth

import (
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var GoogleCmd = &cobra.Command{
	Use:   "google",
	Short: "Login with google",
	Run: func(cmd *cobra.Command, args []string) {
		url := "https://google.com"
		err := exec.Command("xdg-open", url).Start()
		if err != nil {
			cmd.Println("Error opening browser:", err)
			log.Print("Error opening browser:", err)
		}
	},
}
