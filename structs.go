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

// now attach outputUserDetails to the user struct
// u user is a receiver argument
// func (u *user) outputUserDetails() { //<= can use a pointer here as the pointer for more efficiency
func (u user) outputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

// must dereference the pointer to the struct so no extra memory is used
func (u *user) clearUserDetails() {
	u.firstName = ""
	u.lastName = ""
	u.birthDate = ""
}

// creation/constructor function which is a utlity for creating a struct
// by convention precede the struct name w/ new
func newUser(firstName, lastName, birthDate string) *user {
	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userbirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user

	appUser = newUser(userFirstName, userLastName, userbirthDate)

	// ... do something awesome with that gathered data!
	// now using the receiver method w/o any arguments
	appUser.outputUserDetails()
	appUser.clearUserDetails()
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
