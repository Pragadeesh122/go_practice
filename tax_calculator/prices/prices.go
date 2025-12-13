package prices

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strconv"
)

type TaxIncludePriceJob struct{
	TaxRate float64
	Prices []float64
	TaxIncludedPrices map[string]float64
}



func (job *TaxIncludePriceJob) LoadPrices(){

	file, err := os.Open("tax_calculator/prices/prices.txt")

	if err != nil {
		fmt.Println("there was an error reading the file : ", err)
		return
	}

	scanner := bufio.NewScanner(file)
	prices := []float64{}
	for scanner.Scan(){
		converted, err := strconv.ParseFloat(scanner.Text(),64)
		if err != nil{
			fmt.Println("there was an error converting the string to integer, check if the value provided is an integer ", err)
			return
		}
		prices = append(prices, converted)
	}

    err = scanner.Err()

	if err != nil{
		fmt.Println("there was an error scanning the text and the scanner failed ", err)
		file.Close()
		return
	}

	job.Prices = prices

	file.Close()

}

func (job *TaxIncludePriceJob) CalculateTaxIncludedPrices() {
	result := make(map[string]float64)

	for _, value := range job.Prices{
		result[strconv.FormatFloat(value,'f',2,64)] = math.Round(value * (1+job.TaxRate))
	}
	job.TaxIncludedPrices = result
	job.WriteDataToFile(fmt.Sprintf("result_%v.json",job.TaxRate*100))
}

func (job *TaxIncludePriceJob) OutputPrices(){
	for key, value := range job.TaxIncludedPrices{
		fmt.Printf("For the Tax Rate %v The prices before and after are %v and %v\n",job.TaxRate, key, value)
	}
}

func New(taxRate float64) *TaxIncludePriceJob{

	return &TaxIncludePriceJob{
		TaxRate: taxRate,
	}
}


func (job *TaxIncludePriceJob) WriteDataToFile(path string){

	file, err := os.Create(path)

	if err != nil{
		fmt.Println(err)
		return
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(job)

	if err != nil{
		fmt.Println(err)
		file.Close()
		return
	}

	file.Close()
}