package main

import "fmt"

func main() {
	// var names map[string]int // names ' in bir map olduğunu söylüyoruz

	// names = make(map[string]int, 0) // names'i map yapıyoruz ve 0 dan başaltıyoruz

	// names["Ünsal"] = 1 // değer atama key:value
	// names["ünsal"] = 2 // değer atama key:value
	// names["unsal"] = 3 // değer atama key:value

	// fmt.Println(names["meh"])
	//--------------------------------------------------------------------------------
	names := map[string]int{
		"Ünsal": 1,
		"ünsal": 2,
		"unsal": 3,
	}

	delete(names, "unsal") // silmek için

	fmt.Println(names)
}
