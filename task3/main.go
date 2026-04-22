package main

import "fmt"

func MaskCard(cardNumber string) string {
	return fmt.Sprintf("%s **** **** %s", cardNumber[:4], cardNumber[12:])
}

func main() {

}
