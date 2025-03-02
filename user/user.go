package user

import (
	"fmt"
	"time"
)

// globals
type User struct {
	FirstName string
	LastName  string
	BirthDate string
	CreatedAt time.Time
}

// Admin struct which embeds the User struct
type Admin struct {
	Email    string
	Password string
	User
}

// now attach outputUserDetails to the user struct
// u user is a receiver argument
// func (u *user) outputUserDetails() { //<= can use a pointer here as the pointer for more efficiency
func (u User) OutputUserDetails() {
	fmt.Println(u.FirstName, u.LastName, u.BirthDate)
}

// must dereference the pointer to the struct so no extra memory is used
func (u *User) ClearUserDetails() {
	u.FirstName = ""
	u.LastName = ""
	u.BirthDate = ""
}

func NewAdmin(email, password, firstName, lastName, birthDate string) (*Admin, error) {
	//add validation steps
	if email == "" || password == "" || firstName == "" || lastName == "" || birthDate == "" {
		return nil, fmt.Errorf("missing required fields")
	}

	return &Admin{
		Email:    email,
		Password: password,
		User: User{
			FirstName: firstName,
			LastName:  lastName,
			BirthDate: birthDate,
			CreatedAt: time.Now(),
		},
	}, nil
}

// creation/constructor function which is a utlity for creating a struct
// by convention precede the struct name w/ new
func New(firstName, lastName, birthDate string) (*User, error) {
	//add validation steps
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, fmt.Errorf("missing required fields")
	}

	return &User{
		FirstName: firstName,
		LastName:  lastName,
		BirthDate: birthDate,
		CreatedAt: time.Now(),
	}, nil
}
