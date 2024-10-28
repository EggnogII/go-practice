package prices

import (
	"fmt"

	"example.com/majima-tax-program/conversion"
	"example.com/majima-tax-program/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := job.IOManager.ReadLines()

	// Convert string valued input prices to float values
	prices := make([]float64, len(lines))
	prices, err = conversion.StringsToFloats(lines)

	if err != nil {
		fmt.Println(err)
		return
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
	job.TaxIncludedPrices = result
	job.IOManager.WriteResult(job)
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   iom,
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
