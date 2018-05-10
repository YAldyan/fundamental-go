package main

import (
	"fmt"
)

func main() {

	// operator aritmatika
	var value = (((2+6)%3)*4 - 2) / 3

	// operator perbandingan
	var isEqual = (value == 2)

	fmt.Printf("nilai %d (%t) \n", value, isEqual)

	/*
		operator logika

		Template \t digunakan untuk menambahkan indent tabulasi.
		Biasa dimanfaatkan untuk merapikan tampilan output pada
		console
	*/
	var left = false
	var right = true

	var leftAndRight = left && right
	fmt.Printf("left && right \t(%t) \n", leftAndRight)

	var leftOrRight = left || right
	fmt.Printf("left || right \t(%t) \n", leftOrRight)

	var leftReverse = !left
	fmt.Printf("!left \t\t(%t) \n", leftReverse)
}
