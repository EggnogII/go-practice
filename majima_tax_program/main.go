package main

import (
	"fmt"

	"example.com/majima-tax-program/filemanager"
	"example.com/majima-tax-program/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	// Iterate through and associate
	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		//cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		priceJob.Process()
	}

}
