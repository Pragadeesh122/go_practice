package admin


import (
	"fmt"
	"go_practice/user"
	"errors"
)


type Admin struct{
	user.User
	email string
	passowrd string
}


func (admin *Admin) OutputUser(){
	admin.User.OutputUser()
	fmt.Println("email", admin.email)
	
}

func New(user *user.User, email string, password string)(*Admin,error){
	if email == "" || password == ""{
		return nil, errors.New("the email or password should not be empty")
	}
	
	return &Admin{
		*user,
		email,
		password,
	}, nil
}

