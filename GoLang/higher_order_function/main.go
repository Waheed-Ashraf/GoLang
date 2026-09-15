package main

import (
	"fmt"
	"higher_order_function/src"
)

func main() {
	numbers := []int{1, 4, 3, 4, 5}

	mappedSlice := src.MapIntNums(numbers, func(num int) int {
		return num * 2
	})
	fmt.Println(mappedSlice)

	filteredSlice := src.FilterSlice(numbers, func(num int) bool {
		return num > 3
	})
	fmt.Println(filteredSlice)

	reducedValue := src.ReduceSlice(numbers, 1, func(acc, number int) int {
		return acc * number
	})
	fmt.Println(reducedValue)

}
