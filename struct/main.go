package main

import (
	"fmt"
	"go_practice/user"
)



func main(){

	 appUser, err  := user.New( "Pete", "Griffin","02/23/2334")

	 if err != nil{
		fmt.Println("Please enter a full user name")
		return
	 }
	
	appUser.OutputUser()
	appUser.ClearName()
	appUser.OutputUser()
}


