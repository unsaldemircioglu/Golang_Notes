package main

import "fmt"

var x = 10

func main() {
	// var condition = true

	// if condition {
	// 	//block
	// 	fmt.Println(condition)
	// }
	// fmt.Println(condition)

	// test()
	// fmt.Println(x, y)
	fmt.Println(x)
	test()
}
func test() {
	var x = 10 // fucntion scope
	var y = 20
	fmt.Println(x, y)
}
