package prices

type TaxIncludedPriceJob struct {
	TaxRate float64
	InputPrices []float64
	TaxIncludedPrices []float64
}
for _, taxRate := range taxRates {
	var taxIncludedPrices []float64 = make([]float64, len(prices))
	for priceIndex, price := range prices {
		taxIncludedPrices[priceIndex] = price * (1 + taxRate)
	}
	result[taxRate] = taxIncludedPrices
}