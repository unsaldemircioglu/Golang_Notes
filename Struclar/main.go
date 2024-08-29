package main

import "fmt"

func main() {
	// var customer1 = Customer{id: 1, name: "Ünsal Demircioğlu"}

	//--------------------------
	// customer1.ToString()
	// changeName(customer1)
	// customer1.ToString()
	//--------------------------
	// var customer2 = Customer{id: 2, name: "Unsal Demircioğlu", age: 55}

	// customer1.age = 31
	// fmt.Println(customer1)
	// fmt.Println(customer2)
	// var customer1 Customer
	// customer1 = Customer{id: 1, name: "Ünsal Demircioğlu", age: 30}
	//--------------------------
}
func (customer Customer) changeName(newName string) {
	customer.name = newName
}

// func changeName(customer Customer) {
// 	customer.name = "Unsal"
// }

type Customer struct {
	id   int64
	name string
	age  int
}
type Address struct {
	city     string
	district string
}

func (customer Customer) ToString() {
	fmt.Printf("Name %s, Age: %d", customer.name, customer.age)
}
