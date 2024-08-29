package main

import "fmt"

func main() {
	// defer fmt.Println("Hello") // Functionm işlevi bittiğinde en son çalışır(Dosya Kapatılması,Veri Tabanı kapatılıması,fonksiyn çalıştıkan sonra bir şeyleri temizlemek için gibi yerlerde kullanılır).
	// fmt.Println("World")
	//-------------------------------------------------------------------------
	// Out:
	// World
	// Hello
	//-------------------------------------------------------------------------
	// defer fmt.Println("1.Defer")
	// defer fmt.Println("2.Defer")
	// defer fmt.Println("3.Defer")

	// fmt.Println("Main Fonksiyonu")
	//-------------------------------------------------------------------------
	// var condition = true
	// defer fmt.Println("1.Defer")
	// if condition {
	// 	return
	// }
	// defer fmt.Println("2.Defer")
	//-------------------------------------------------------------------------
	// x := 10
	// y := 20
	// defer fmt.Println("x:", x)
	// defer fmt.Println("y:", y)
	//-------------------------------------------------------------------------
	var condition = true
	defer cleanup() // bunu yaptığımız için program panic atsa bile çalışır çünkü bu işlem bitince çalışıyor.
	if condition {
		panic("An error occurred") // Programı sonladırıp altındaki işlemi gerçekleştirmiyor
	}
}

func cleanup() {
	fmt.Println("Clenup worked...")
}
