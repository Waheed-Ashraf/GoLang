package main

import "fmt"

// output questions

// 1 - 1 , 100

// 2- 1

// 3- true

// 4- 99

// 5- 1 2 3

//

type Engine struct{ HP int }
type Car struct{ *Engine }

func main() {
	c := Car{}
	fmt.Println(c.Engine == nil)
}
