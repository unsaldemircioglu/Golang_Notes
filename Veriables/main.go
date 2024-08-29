package main

import (
	"fmt"
)

func main() {
	// var productName string
	// var quantity int
	// var discount float32
	// var isInStock bool

	// productName = "Golang Lesson"
	// quantity = 10
	// discount = 9.37
	// isInStock = true

	// fmt.Println(productName)
	// fmt.Println(quantity)
	// fmt.Println(discount)
	// fmt.Println(isInStock)

	// fmt.Println(productName, reflect.TypeOf(productName))
	// fmt.Println(quantity, reflect.TypeOf(quantity))
	// fmt.Println(discount, reflect.TypeOf(discount))
	// fmt.Println(isInStock, reflect.TypeOf(isInStock))

	// var productName = "Golang Lesson"
	// fmt.Println(productName, reflect.TypeOf(productName))

	// productName := "Golang lesson"
	// fmt.Println(productName)

	// quentity := 10
	// fmt.Println(quentity)

	var productName string
	var quantity int
	var discount float32
	var isInStock bool

	productName = "Golang Lesson"
	quantity = 10
	discount = 9.37
	isInStock = true

	// fmt.Println("Product Name: ", productName, "Quantity: ", quantity, "Discount: ", discount, "IsInStock:", isInStock)
	fmt.Printf("Product Name: %s, Quantity: %d , Discount: %f, IsInStock: %t", productName, quantity, discount, isInStock)

}
