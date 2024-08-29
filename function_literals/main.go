package main

import (
	"fmt"
)

func main() {
	f1()
	//-------------------------------------------
	// add := func(x int, y int) int {
	// 	return x + y
	// }
	// fmt.Println(reflect.TypeOf(add))
	// fmt.Println(add(3, 5))
	//-------------------------------------------
	// func(x int, y int) {
	// 	fmt.Println(x + y)
	// }(3, 5) // Bu şekilde fonksiyon başka bir fonksiyonun içinde çalıştırabiliz ve değerleri bu şeklde gönderebiliz
	//-------------------------------------------
	// func() {
	// 	fmt.Println("f2")
	// }() // Bu şekilde fonksiyon başka bir fonksiyonun içinde çalıştırabiliz
}
func f1() {
	fmt.Println("f1")
}
func calculator(x int, y int, f1 func(int, int) int, f2 func(int, int) int) (int, int) {
	return f1(x, y), f2(x, y)
}
