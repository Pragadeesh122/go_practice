package main

import (
	"fmt"
	// "os"
	// "strconv"
	// "strings"
	"go_practice/utils"
)


func main(){
	// var investment int32;
	// var interest float64;
	// var years float64;
	// fmt.Println("Go Application")


	// value := 23434234234234234

	// fileName := "newfile.txt"


	num := 43

	numAddr := &num

	fmt.Println(numAddr)


	// os.WriteFile(fileName,[]byte(strconv.Itoa(value)),0644)

	// fileContent, err := os.ReadFile(fileName)
	// if err != nil{
	// 	fmt.Println("We ran into an error", err)
	// }

	date := utils.GetDate()

	fmt.Println(date)



	// file  := string(fileContent)

	
	// fmt.Printf(`The vale read from the file : %v`,file)
	// fmt.Print("Enter the Investment Amount : ")
	// fmt.Scan(&investment)
	// fmt.Print("Enter the rate of interest : ")
	// fmt.Scan(&interest)
	// fmt.Print("Enter the years to invest : ")
	// fmt.Scan(&years)

	

	// total_returns := float64(investment) * math.Pow((1+interest/100),(years))

	// fmt.Println("Total Return with Captial and Inertest :", total_returns)
	// fmt.Println("An Exmaple funcition below :")
	// next()
	app()
	// res,err := bank_calculator()

	// if err != nil{
	// 	panic("We can't continue because of errors")
	// }

	// fmt.Print("The balance was written successfully :",res)
	
	// 
	
}


// func twoReturns(value1,value2 int)(int,string) {
// 	return value1*value2, fmt.Sprintf("The total value %v",(value1*value2+9923))
// }