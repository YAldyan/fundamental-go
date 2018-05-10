package main

import (
	"fundamental-golang/library"
	"fmt"
)

func main() {
	library.SayHelloX("Yuhariz Aldyan")
	// library.introduce("ethan")

	/*
		Penggunaan Public & Private Pada Struct Dan Propertinya
	*/
	var s1 = library.Student{"ethan", 21}
	fmt.Println("name ", s1.Name)
	fmt.Println("grade", s1.Grade)

	/*
		Mengakses Properti Dalam File Yang Package-nya Sama
	*/
	sayHello("ethan")

	/*
		fungsi init
	*/
	fmt.Printf("Name  : %s\n", library.Student.Name)
	fmt.Printf("Grade : %d\n", library.Student.Grade)
}
