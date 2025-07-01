package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	var appUser *user.User
	appUser, err := user.NewUser("John", "Doe", "1990-01-01")

	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}
	fmt.Println(appUser.String())
	appUser.ClearUserName()
	fmt.Println(appUser.String())
}
