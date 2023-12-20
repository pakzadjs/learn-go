package main

import "fmt"
import "os"

func writeBalanceToFile(balance float64)  {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

func main() {
	var accountBalance float64 = 1000
	fmt.Println("Welcome to Go Bank!")

	// for i := 0; i < 200; i++ {
	for { // Infinite loop

		fmt.Println("Wat do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("your choice:")
		fmt.Scan(&choice)

		// wantsCheckBalance := choice == 1

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Your deposit: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				// return
				continue
			}

			accountBalance += depositAmount // accountBalance = accountBalance + depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			var withdrawAmount float64
			fmt.Print("Withdraw amount: ")
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our Bank")
			return
			// break
		}

		// if choice == 1 {
		// 	fmt.Println("Your balance is:", accountBalance)
		// } else if choice == 2 {
		// 	var depositAmount float64
		// 	fmt.Print("Your deposit: ")
		// 	fmt.Scan(&depositAmount)

		// 	if depositAmount <= 0 {
		// 		fmt.Println("Invalid amount. Must be greater than 0.")
		// 		// return
		// 		continue
		// 	}

		// 	accountBalance += depositAmount // accountBalance = accountBalance + depositAmount

		// 	fmt.Println("Balance updated! New amount:", accountBalance)
		// } else if choice == 3 {
		// 	var withdrawAmount float64
		// 	fmt.Print("Withdraw amount: ")
		// 	fmt.Scan(&withdrawAmount)

		// 	if withdrawAmount <= 0 {
		// 		fmt.Println("Invalid amount. Must be greater than 0.")
		// 		continue
		// 	}
		// 	if withdrawAmount > accountBalance {
		// 		fmt.Println("Invalid amount. You can't withdraw more than you have.")
		// 		continue
		// 	}

		// 	accountBalance -= withdrawAmount
		// 	fmt.Println("Balance updated! New amount:", accountBalance)
		// } else {
		// 	fmt.Println("Goodbye!")
		// 	// return
		// 	break
		// }
	}

	// fmt.Println("Thanks for choosing our Bank")
}
