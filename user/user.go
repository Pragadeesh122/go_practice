package user

import (
	"errors"
	"fmt"
	"time"
)


type User struct {
	firstName string
	lastName string
	birthDate string 
	createdAt time.Time 
}



func (user *User) OutputUser(){
	fmt.Println("First Name", user.firstName)
	fmt.Println("Last Name", user.lastName)
	fmt.Println("BirthDate ", user.birthDate)
	fmt.Println("CreatedAt ", user.createdAt)
}


func (user *User) ClearName(){
	user.firstName = ""
	user.lastName = ""
}


func New(firstName string, lastName string, birthDate string) (*User, error){

	if(firstName == "" || lastName == "" || birthDate == ""){
		return  nil, errors.New("Hds")
	}


	return &User{
		firstName : firstName,
		lastName : lastName,
		birthDate : birthDate,
		createdAt : time.Now(),
	}, nil
}

