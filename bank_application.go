package main

import (
	"fmt"
	"os"
	"errors"
	"strconv"
)


func bank_calculator() (int,error){
	var balance int
	fmt.Print("Enter the balance :")
	_ , err := fmt.Scan(&balance)

	if(err != nil){
		return 0, errors.New("the balance value provided had an error")
	}

	fileErr := os.WriteFile("balance.txt",[]byte(strconv.Itoa(balance)),0644)

	if fileErr != nil{
		return 0, errors.New("the balance cannot be read from the file")
	}

	return 1, nil


}