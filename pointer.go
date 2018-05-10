package main

import (
	"fmt"
)

func main() {

	// inisialisasi pointer
	// var number *int
	// var name *string

	/*
		penerapan pointer
	*/
	var numberA int = 4

	/*
		 	menunjuk ke alamat/reference dari NumberA
			sehingga nilainya pasti sama
	*/
	var numberB *int = &numberA

	fmt.Println("numberA (value)   :", numberA)  // 4
	fmt.Println("numberA (address) :", &numberA) // 0xc20800a220

	fmt.Println("numberB (value)   :", *numberB) // 4
	fmt.Println("numberB (address) :", numberB)  // 0xc20800a220

	/*
		Efek Perubahan Nilai Pointer
	*/
	numberA = 5

	fmt.Println("numberA (value)   :", numberA)
	fmt.Println("numberA (address) :", &numberA)
	fmt.Println("numberB (value)   :", *numberB)
	fmt.Println("numberB (address) :", numberB)

	/*
		Parameter Pointer
	*/
	var number = 4
	fmt.Println("before :", number) // 4

	change(&number, 10)
	fmt.Println("after  :", number) // 10
}

func change(original *int, value int) {
	*original = value
}
