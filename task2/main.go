package main

import "fmt"

func main() {
	var productName string = "iPhone 15 Pro"
	var brand string = "Apple"
	var price int64 = 12990000
	var inStock bool = true

	monthlyPayment := price / 12

	fmt.Println(" Alifshop ")
	fmt.Printf("Товар: %s\n", productName)
	fmt.Printf("Бренд: %s\n", brand)
	fmt.Printf("Цена: %d сум\n", price)
	fmt.Printf("В наличии: %t\n", inStock) // %t для вывода true/false
	fmt.Printf("Рассрочка: 12 мес -> %d сум/мес\n", monthlyPayment)
}
