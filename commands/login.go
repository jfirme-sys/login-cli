package commands

import (
	"envsecret-cli/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var LoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in to the API",
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")

		token, err := utils.Login(username, password)
		if err != nil {
			fmt.Println("Login failed:", err)
			return
		}

		fmt.Println("Login successful, token:", token)
	},
}

func init() {
	LoginCmd.Flags().String("username", "", "Username for login")
	LoginCmd.Flags().String("password", "", "Password for login")
	LoginCmd.MarkFlagRequired("username")
	LoginCmd.MarkFlagRequired("password")
}
