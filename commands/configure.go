package commands

import (
	"envsecret-cli/auth"
	"fmt"
)

func Configure() {
	email := "julianofirme@gmail.com"
	password := "Pass1234"
	token, err := auth.Authenticate(email, password)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Token:", token)
}
