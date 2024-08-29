package main

import (
	"fmt"
	"time"
)

type Product struct {
	Id   int64
	Name string
}

var productChannel = make(chan Product)

func main() {
	select {
	case productDetail := <-ProductChannel:
		fmt.Println("Ürün Detaylarıo")
	}
	fmt.Println("gr")
}
func getProductDeatilsFromExternalService(productId int64) {
	time.Sleep(10 * time.Second)
	productChannel <- Product{
		Id:   productId,
		Name: "Çamaşır Makinesi",
	}
}
