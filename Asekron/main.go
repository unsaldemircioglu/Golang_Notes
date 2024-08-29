package main

import (
	"fmt"
	"sync"
)

func main() {
	// go f1() //
	// go f2()
	// fmt.Println("End Of The Main")
	// time.Sleep(3 * time.Second)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("f1")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("f2")
	}()
	wg.Wait() // Blocked
	fmt.Println("End Of The Main")
}

// func f1() {
// 	fmt.Println("f1")
// }
// func f2() {
// 	fmt.Println("f2")
// }
