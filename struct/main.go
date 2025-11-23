package main

import (
	"fmt"
	"go_practice/user"
	"go_practice/admin"
)



func main(){

	 appUser, err  := user.New( "Pete", "Griffin","02/23/2334")

	 if err != nil{
		fmt.Println("Please enter a full user name")
		return
	 }
	
	 admin, errAdmin := admin.New(appUser,"asdas@gmail.com","dasdasdasd")

	 if errAdmin != nil{
		fmt.Println(errAdmin)
		return
	 }
	admin.OutputUser()
	admin.ClearName()
	admin.OutputUser()
}


