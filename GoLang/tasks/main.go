package main

import "fmt"

// task 1 memoize factorial function [closure , map ]  {4:24 , 5:120 , 6:720 , 7:5040 , 8:40320 , 9:362880 , 10:3628800}
func memoizeFactorial() func(int) int {
	memo := make(map[int]int)

	var factorial func(int) int

	factorial = func(n int) int {
		if val, ok := memo[n]; ok {
			return val
		}

		if n == 0 {
			memo[n] = 1
			return 1
		}

		memo[n] = n * factorial(n-1)
		return memo[n]
	}

	return factorial
}

// check if string is palindrome or not
func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

// check if two strings are anagrams of each other
// func areAnagrams(s1, s2 string) bool {
// 	if len(s1) != len(s2) {
// 		return false
// 	}
// 	for i := 0; i < len(s1); i++ {
// 		for j := 0; j < len(s2); j++ {
// 			if s1[i] == s2[j] {
// 				break
// 			}
// 			return false
// 		}

//		}
//		return true
//	}
func areAnagrams(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	counts := make(map[rune]int)

	for _, char := range s1 {
		counts[char]++
	}

	for _, char := range s2 {
		if counts[char] == 0 {
			return false
		}
		counts[char]--
	}

	return true
}

// Task 4: Return the average of any number of values.
func average(numbers ...float64) (float64, error) {
	if len(numbers) == 0 {
		return 0, fmt.Errorf("provide at least one number")
	}

	total := 0.0
	for _, number := range numbers {
		total += number
	}

	return total / float64(len(numbers)), nil
}

// Task 5: Simple calculator with error handling.
func calculate(a, b float64, operation string) (float64, error) {
	switch operation {
	case "sum":
		return a + b, nil
	case "diff":
		return a - b, nil
	case "mul":
		return a * b, nil
	case "div":
		if b == 0 {
			return 0, fmt.Errorf("cannot divide by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("unknown operation: %s", operation)
	}
}

// Task 6: Functions summary.
// Parameters provide inputs; return values provide outputs.
// A function can return multiple values, such as a result and an error.
// Variadic parameters (...float64) accept any number of arguments.
// Closures retain access to surrounding variables, such as memo.
// Recursion means a function calls itself, such as factorial(n-1).
// Check err != nil to handle errors; nil means no error occurred.
func main() {
	f := memoizeFactorial()
	fmt.Println(f(4))                    // 24
	fmt.Println(f(5))                    // 120
	fmt.Println(isPalindrome("racecar")) // true
	fmt.Println(isPalindrome("hello"))   // false

	fmt.Println(areAnagrams("listen", "silent")) // true

	avg, err := average(10, 20, 30)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Average:", avg) // 20
	}

	for _, operation := range []string{"sum", "diff", "mul", "div"} {
		result, err := calculate(10, 2, operation)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(operation, result)
		}
	}

	// Examples of invalid inputs.
	if _, err := average(); err != nil {
		fmt.Println(err)
	}
	if _, err := calculate(10, 0, "div"); err != nil {
		fmt.Println(err)
	}
	if _, err := calculate(10, 2, "unknown"); err != nil {
		fmt.Println(err)
	}
}
