package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	var appUser *user.User
	appUser, err := user.New("John", "Doe", "1990-01-01")

	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}
	admin := user.NewAdmin("John", "Doe", "1990-01-01", "asdas", "sdfsd")
	fmt.Println(appUser.String())
	appUser.ClearUserName()
	fmt.Println(appUser.String())
	fmt.Println(admin.String())
}
