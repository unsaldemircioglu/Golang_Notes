package main

import (
	"fmt"
)

func main() {
	// data := f1()
	// fmt.Println(data)
	//-------------------------
	myChannel := make(chan string)
	go func() {
		myChannel <- "Hello World"
	}()
	message := <-myChannel

	fmt.Println(message)
}
func f1() int {
	return 1
}
