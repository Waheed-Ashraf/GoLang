package main

import (
	"errors"
	"fmt"
)

func calcBalance(balance, withdraw int) (int, error) {
	if withdraw > balance {
		return 0, errors.New("insufficient funds")
	}
	return balance - withdraw, nil
}

func main() {

	balance, err := calcBalance(1000, 2000)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Remaining balance:", balance)
}

// type months int
// const(
// 	jan = iota + 1
// 	fab
// 	jun
// )
// func main() {
//    fmt.Println("Hello, World!")
//    var price = 1000
//    var taxRetio = .14
//    var priceAfterRetio  =float64(price) * taxRetio
//    fmt.Println("Price after tax ratio is: ", priceAfterRetio)
// }
