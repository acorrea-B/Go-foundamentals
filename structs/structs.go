package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func (u *User) String() string {
	return fmt.Sprintf("Name: %s %s, Birthdate: %s, Created At: %s", u.firstName, u.lastName, u.birthdate, u.createdAt.Format(time.RFC3339))
}

func (u *User) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func main() {
	var appUser User
	appUser = User{
		firstName: "John",
		lastName:  "Doe",
		birthdate: "1990-01-01",
		createdAt: time.Now(),
	}
	appUser.clearUserName()
	fmt.Println(appUser.String())
}
