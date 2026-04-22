package main

import "fmt"

func ExampleMaskCard_success() {
	result := MaskCard("4111111111111111")
	fmt.Println(result)

}

func ExampleMaskCard_panic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Паника успешно поймана!")
		}
	}()

	MaskCard("123456")

}
