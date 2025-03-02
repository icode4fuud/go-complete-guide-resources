package main

import (
	"fmt"
	"time"
)

// globals
type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userbirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser := user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userbirthDate,
		createdAt: time.Now(),
	}

	// ... do something awesome with that gathered data!
	// now passing a pointer instead of an unnecessary copy
	outputUserDetails(&appUser)
}

func outputUserDetails(u *user) { //<= can use a pointer here as the pointer for more efficiency
	//dereference to perform data mutation
	(*u).firstName = "John"
	(*u).lastName = "Doe"
	(*u).birthDate = "01/01/1970"

	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
