package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// var x int
	// var y float32
	// var pointer1 *int

	// fmt.Println(x)
	// fmt.Println(y)
	// fmt.Println(pointer1)
	//------------------------------------------------
	// var pointer1 *int

	// if pointer1 == nil {
	// 	fmt.Println("Pointer Herhangi bir yeri göstermiyor.")
	// }
	//------------------------------------------------
	// fileData, _ := os.ReadFile("sample.txt") // Eğer eror'ları kontrol etmeyecek isek err yerine _ yazabiliriz(Tavsiye Edilmez)
	fileData, err := os.ReadFile("sample.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fileData)
	}

}
func divide(x int, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Payda 0 Olamaz!")
	}
	return x / y, nil
}
