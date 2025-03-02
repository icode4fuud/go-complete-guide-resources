package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userbirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User

	appUser, err := user.New(userFirstName, userLastName, userbirthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	// ... do something awesome with that gathered data!
	// now using the receiver method w/o any arguments
	appUser.OutputUserDetails()
	appUser.ClearUserDetails()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
