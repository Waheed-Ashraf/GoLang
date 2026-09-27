package main

import (
	"fmt"
)

// type IShape interface {
// 	Area() float64
// 	Perimeter() float64
// }

// type Circle struct {
// 	Radius float64
// }

// type Square struct {
// 	Side float64
// }

// func (c Circle) Area() float64 {
// 	return math.Pi * math.Pow(c.Radius, 2)
// }

// func (c Circle) Perimeter() float64 {
// 	return math.Pi * c.Radius * 2
// }

// func (sq Square) Area() float64 {
// 	return math.Pow(sq.Side, 2)
// }

// func (sq Square) Perimeter() float64 {
// 	return sq.Side * 4
// }

// func ShowDetails(s IShape) {
// 	fmt.Println("Area = ", s.Area())
// 	fmt.Println("Perimeter = ", s.Perimeter())
// }

// // calc toatal area of all shapes
// func calcTotalArea(shapes []IShape) float64 {
// 	var totalArea float64
// 	for _, shape := range shapes {
// 		totalArea += shape.Area()
// 	}
// 	return totalArea
// }

// // calc largest area of all shapes
// func calcLargestArea(shapes []IShape) float64 {
// 	var largestArea float64
// 	for _, shape := range shapes {
// 		if shape.Area() > largestArea {
// 			largestArea = shape.Area()
// 		}
// 	}
// 	return largestArea
// }

// func main() {
// 	c := Circle{Radius: 8}
// 	d := Circle{Radius: 20}

// 	sq := Square{10}
// 	sq2 := Square{20}

// 	ShowDetails(c)
// 	ShowDetails(d)
// 	ShowDetails(sq)
// 	ShowDetails(sq2)
// 	fmt.Println("Total Area = ", calcTotalArea([]IShape{c, d, sq, sq2}))
// 	fmt.Println("Largest Area = ", calcLargestArea([]IShape{c, d, sq, sq2}))
// }

// type Speaker interface {
// 	Speak() string
// }
// type Dog struct{}

// func (d *Dog) Speak() string {
// 	return "Woof"
// }

func main() {
	var num interface{} = int32(10)
	var num2 interface{} = int64(10)
	fmt.Println(num == num2)
}
