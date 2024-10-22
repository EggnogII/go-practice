package prices

import "fmt"

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job TaxIncludedPriceJob) Process() {
	result := make(map[string]float64)
	for _, price := range job.InputPrices {
		resultStr := fmt.Sprintf("%0.2f", price)
		result[resultStr] = price * (1 + job.TaxRate)
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
