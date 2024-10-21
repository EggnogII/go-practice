package main

import (
	"fmt"
)

func main() {
	num := multiply(3, 2)
	fmt.Println(num)
	calculate_profit()
}

func multiply(num int, second_num int) int {
	// also (num, second_num int) as param is acceptable
	mult := num * second_num
	return mult
}

func get_input_number(flavor_text string) int {
	var number int = 0
	fmt.Print(flavor_text)
	fmt.Scan(&number)
	return number
}

func get_tax_rate() float64 {
	var tax_rate float64 = 0.0
	fmt.Print("Enter your tax rate: ")
	fmt.Scan(&tax_rate)
	return tax_rate
}
func calculate_profit() {
	var revenue_amount int = get_input_number("Enter Revenue: ")
	var expenses int = get_input_number("Enter your Expenses: ")
	var tax_rate float64 = get_tax_rate()

	earnings_before_tax := revenue_amount - expenses
	fmt.Print("\nEBT: ")
	fmt.Print(earnings_before_tax)

	earnings_after_tax := calculate_earnings_after_tax(tax_rate, earnings_before_tax)

	fmt.Print("\nEAT: ")
	fmt.Print(earnings_after_tax)

	ratio := calculate_ratio(float64(earnings_before_tax), float64(revenue_amount))

	fmt.Print("\nRatio: ")
	fmt.Print(ratio)
}

func calculate_ratio(earnings_before_tax float64, revenue_amount float64) float64 {
	ratio := earnings_before_tax / revenue_amount
	return ratio
}

func calculate_earnings_after_tax(tax_rate float64, earnings_before_tax int) float64 {
	true_rate := tax_rate / 100.0
	earnings_after_tax := float64(earnings_before_tax) * true_rate
	return earnings_after_tax
}
