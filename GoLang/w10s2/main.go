package main

import (
	"fmt"
)

func withdraw(balance, amount int) (int, error) {
	if amount > balance {
		return balance, fmt.Errorf(
			"balance value is %d and amount is %d",
			balance, amount,
		)
	}
	return balance - amount, nil
}
func main() {

	value, err := withdraw(100, 150)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("New balance:", value)
	}
}
