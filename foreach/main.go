package main

import "fmt"

func main() {
	//-----------------------------------------------------
	// var numbers = []int{1, 2, 3, 4}

	// for index := 0; index < len(numbers); index++ {
	// 	fmt.Println(numbers[index])
	// }

	// for index, value := range numbers {
	// 	fmt.Println(index, value)
	// }

	// for _, value := range numbers { // Eğer  _ ifadesini kullanırsak değişkeni kullanmasak bile hata vermez.
	// 	fmt.Println(value)
	// }
	//-----------------------------------------------------
	// var language = "Golang"

	// for _, charecter := range language {
	// 	fmt.Printf("Charecter %c\n", charecter)
	// }
	//-----------------------------------------------------
	var names = map[string]int{
		"Ünsal": 10,
		"ünsal": 20,
		"unsal": 30,
	}
	for key, value := range names {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}
