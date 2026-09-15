package main

import "fmt"

// q1 output is ==> first 30 second 30
// q2 print numbers from 10 to 1 without using loop

func printNumbers(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	printNumbers(n - 1)
}

// q3 output is ==> [2,4,6]

func main() {
	printNumbers(10)
}
