package main

import "fmt"

func main() {
	fmt.Println("This is a standard string print.")
	var number float64 = 34.567890
	// These are some ways to format string during print
	fmt.Println("This is a float:", number)
	fmt.Printf("This is a float printed via format: %v.\n", number)
	fmt.Printf("This is a float printed via mod format: %.2f \n", number)

	// This is how you make a formatted string for later printing
	message := fmt.Sprintf("Some number %0.2f\n", number)
	fmt.Print(message)

	// This is how you do multiline
	fmt.Printf(`This is some number in 
	a multiline string that's being printed out immediately: %0.3f
	for kicks I even modified the number of decimal points in the float response.
	`, number)
}
