package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)




func app(){

	revenue, revenueErr := getRevenue()

	if revenueErr != nil{
		fmt.Println("Error", revenueErr)
		return
	}

	expense, expenseErr := getExpense()

	if expenseErr != nil{
		fmt.Println("Error", expenseErr)
		return
	}
	
	taxRate , taxRateErr := getTaxRate()

	if taxRateErr != nil{
		fmt.Println("Error", taxRateErr)
		return
	}


	_,_,_,calculateEarningsErr := calculateEarnings(revenue,expense,taxRate)

	if calculateEarningsErr != nil{
		fmt.Println("Error", calculateEarningsErr)
		return
	}


	
}



func getRevenue() ( int,  error){


	for{

		var revenue int 
	
		fmt.Print("Revenue : ")
		_,err:=fmt.Scan(&revenue)
	
		if err != nil {
			fmt.Println("there was an error getting the revenue from the user, Retrying :")
			continue
		}
	
		if revenue <= 0{
			fmt.Println("the revenue should be greater than 0, Please enter a valid revenue")
			continue
		}
	
		return revenue, nil
	}
}


func getExpense() ( int,  error){



	for{
		
		var expense int
	
		fmt.Print("Expenses : ")
		_, err := fmt.Scan(&expense)
	
		if err != nil {
			fmt.Println("there was an error getting the expense from the user, Retrying :")
			continue
		}
	
		if expense < 0{
			fmt.Println("the expense should be greater than or equal 0, Please enter a valid expense value")
			continue
		}
	
		return expense , nil
	}

}


func getTaxRate() ( float64,  error) {

	for{

		var taxRate float64

		fmt.Print("Tax Rate : ")
		_, err := fmt.Scan(&taxRate)


		if err != nil {
			fmt.Println("there was an error getting the tax rate  from the user, Retrying :")
			continue
		}

		if taxRate < 0{
			fmt.Println("the tax rate should be greater than or equal 0, Please enter a valid Tax Rate")
			continue
		}
		return taxRate, nil
	}

}


func calculateEarnings(revenue int, expense int, taxRate float64) ( int,  float64,  float64,  error ){


	earningsBeforeTax :=  revenue - expense
	profit := (float64(revenue) - (float64(revenue) * (taxRate/100))) - float64(expense)

	profitRatio := float64(earningsBeforeTax )/ profit


	data := []string{
			fmt.Sprintf("Earnings before Tax : %v\n", earningsBeforeTax),
			fmt.Sprintf("Profit : %.2f\n", profit),
			fmt.Sprintf("Earnings/Profit Ratio : %.2f\n", profitRatio),
	} 

	err := os.WriteFile("revenue_calcualtions.txt",[]byte(strings.Join(data, "\n")),0644)

	if err != nil{
		return 0,0,0,errors.New("there was an error writing the results to the output file")
	}
	

	return earningsBeforeTax,profit, profitRatio, nil

}