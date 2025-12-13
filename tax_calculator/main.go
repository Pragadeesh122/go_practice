package main

import (
	"go_practice/tax_calculator/prices"
)

func main(){


	taxRates := []float64{.1,0.08,.2,.4}


	for _, taxRate := range taxRates{

		taxPrices := prices.New(taxRate)
		taxPrices.LoadPrices()
		taxPrices.CalculateTaxIncludedPrices()
		taxPrices.OutputPrices()
	}

	

}
