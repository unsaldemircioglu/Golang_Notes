package main

import "fmt"

func main() {
	// Fixed
	// var names [3]string
	// names[0] = "Ünsal"
	// names[1] = "Demircioglu"
	// names[2] = "Unsal"

	// fmt.Println(names)

	var names = [4]string{"Ünsal", "Unsal", "Demircioğlu"}
	names[1] = "unsal"
	fmt.Println(names[0], names[2])
	fmt.Println(names[0:2])

	// var names = []string{"Ünsal","Unsal","Demircioğlu"}
	// names = append(names,"Unsal")
	// fmt.Println(names)

}
