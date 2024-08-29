package main

type Book struct {
	desi int
}

func (book *Book) ShippingCost() int {
	return 5 + book.desi*2
}
func main() {
	// book1 := &Book{desi: 10}

	// fmt.Println(book1.ShippingCost())

}
func calculateTotalShippingCost(books []Book) {
	total := 0
	for _, book := range books {
		total += book.ShippingCost()
	}
	return total
}
