package main

import "fmt"

func main() {
	var operate_bank bool = true
	var balance float64 = 0.0
	var amount float64 = 0.0
	var choice int = 0
	for operate_bank {
		fmt.Println("Welcome to go bank.")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check my balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)
		fmt.Println("Your choice", choice)

		if choice == 1 {
			fmt.Println(balance)
		} else if choice == 2 {
			fmt.Print("Enter an amount: ")
			fmt.Scan(&amount)
			balance := balance + amount
			fmt.Println("New balance: ", balance)
		} else if choice == 3 {
			fmt.Print("Enter an amount to withdraw: ")
			fmt.Scan(&amount)
			balance := balance - amount
			fmt.Println("Balance after withdrawal: ", balance)
		} else if choice == 4 {
			operate_bank = false
		}

	}

}
