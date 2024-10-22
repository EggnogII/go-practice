package main

import (
	"example.com/majima-tax-program/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	// Iterate through and associate
	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.Process()
	}

}
