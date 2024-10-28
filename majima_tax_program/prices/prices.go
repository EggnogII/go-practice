package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) LoadData() {
	// Open the file
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("An error occured!")
		fmt.Println(err)
		return
	}

	// Read the contents
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Return an error if there is an error during read
	err = scanner.Err()
	if err != nil {
		fmt.Println("Reading content failed")
		fmt.Println(err)
		file.Close()
		return
	}

	prices := make([]float64, len(lines))
	for lineIndex, line := range lines {
		floatPrice, err := strconv.ParseFloat(line, 64)

		if err != nil {
			fmt.Println("Converting price to float failed")
			fmt.Println(err)
			file.Close()
			return
		}

		prices[lineIndex] = floatPrice
	}

	job.InputPrices = prices

}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		resultStr := fmt.Sprintf("%0.2f", price)
		result[resultStr] = fmt.Sprintf("%0.2f", taxIncludedPrice)
	}
	fmt.Println(result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}

/*
for _, taxRate := range taxRates {
	var taxIncludedPrices []float64 = make([]float64, len(prices))
	for priceIndex, price := range prices {
		taxIncludedPrices[priceIndex] = price * (1 + taxRate)
	}
	result[taxRate] = taxIncludedPrices
}
*/
