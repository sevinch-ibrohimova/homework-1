package main

import "fmt"

func main() {
	sender := "Sherali"
	receiver := "Alisher"
	amount := 500000

	commission := amount * 1 / 100
	total := amount + commission

	fmt.Printf("Отправитель: %s\n", sender)
	fmt.Printf("Получатель: %s\n", receiver)
	fmt.Printf("Сумма: %d сум\n", amount)
	fmt.Printf("Комиссия: %d сум\n", commission)
	fmt.Printf("Итого: %d сум\n", total)
}
