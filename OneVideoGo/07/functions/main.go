package main

import "fmt"

func main() {
	/*total := add(5, 10)
	fmt.Println(total)*/

	// total, diffrence, multiply := calculations(10, 20)
	// fmt.Println(total, diffrence, multiply)

	// var numbers = [5]int{10, 20, 2, 3, 5}
	// fmt.Println(sum(numbers))

	fmt.Println(sum(3, 4, 5))

}

func sum(x int, y int, z int) int {
	return x + y + z
}

func sum(x int, y int, z int) int {
	return x + y + z
}

// func sum(numbers [5]int) int {
// 	sum := 0

// 	for _, value := range numbers {
// 		sum += value
// 	}
// 	return sum
// }

// func calculations(x int, y int) (int, int, int) {
// 	return x + y, x - y, x * y
// }

// func add(x int, y int) {
// 	fmt.Println(x + y)
// }

// func add(x int, y int) int {
// 	return x + y
// }
