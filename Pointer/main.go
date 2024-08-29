package main

import (
	"fmt"
)

func main() {
	// var a int
	// var p *int

	// a = 10

	// p = &a
	// fmt.Println(a)
	// *p = 20

	// fmt.Println(a)
	// fmt.Println(p)
	//-------------------------------------------
	// var a = 10
	// var b int
	// var p *int

	// b = a
	// p = &a
	// *p = 20
	// fmt.Println(a, b)
	//------------------------------------------------------
	// var a int = 10
	// fmt.Println(a)
	// add12pointer(&a)
	// add12(a)
	// fmt.Println(a
	//------------------------------------------------------
	var numbers = []int{1, 2, 3}
	fmt.Println(numbers)
	changeValues(numbers)
	fmt.Println(numbers)
}
func changeValues(number []int) {
	number[0] = 1000
}
func add12(x int) {
	x = x + 12
}
func add12pointer(x *int) {
	*x = *x + 12
}
