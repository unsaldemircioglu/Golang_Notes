package main

import "fmt"

func main() {
	// index := 1
	// for index <= 10 {
	// 	fmt.Println(index)
	// 	index = index + 1 // index++
	// }
	//-----------------------------------------
	// total := 0
	// for index := 1; index <= 10; index++ {
	// 	// fmt.Println(index)
	// 	total = total + index
	// }
	// fmt.Println(total)

	// index := 0
	// var names = [3]string{"unsal", "Ünsal", "ünsal"}
	// for index < 3 {
	// 	fmt.Println(names[index])
	// 	index++
	// }
	for index := 0; index <= 10; index++ {
		if index == 2 || index == 5 {
			continue
		}
		fmt.Println(index)
	}

}
