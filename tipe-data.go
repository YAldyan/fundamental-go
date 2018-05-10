package main

import (
	"fmt"
)

func main() {
	/*
		Deklarasi Varibel dengan Tipe Data
	*/

	var firstname string = "Yuhariz"

	var lastname string
	lastname = "Aldyan"

	fmt.Printf(" Halo %s %s !! \n", firstname, lastname)

	// menggunakan fungsi Println tidak perlu \n,
	// karena otomatis melakukan enter

	fmt.Println(" Halo", firstname, lastname, "!!")

	/*
		Deklarasi Varibel tanpa Tipe Data
		Tipe Data disesuaikan dengan nilai
		yang di inisialisasi di pada vari-
		bel tersebut.
	*/
	var firstnameX = "Yuhariz"
	lastnameX := "Aldyan"

	fmt.Printf(" Halo %s %s !! \n", firstnameX, lastnameX)

	/*
		Deklarasi Multi Variabel

		var fourth, fifth, sixth string = "empat", "lima", "enam"
		seventh, eight, ninth := "tujuh", "delapan", "sembilan"
		one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"
	*/

	/*
			Variabel underscore
			berperan sebagai keranjang sampah
			digunakan untuk menampung nilai
			variabel yang tidak dipakai.

			_ = "belajar Golang"
		 	= "Golang itu mudah"
			name, _ := "john", "wick"
	*/

	/*
		Deklarasi Variabel dengan Keyword New
		untuk mencetak data pointer
	*/

	nameX := new(string)

	fmt.Println(nameX)  // 0x20818a220
	fmt.Println(*nameX) // ""

	/*
		Tipe Data Numerik Non Desimal

		uint, tipe data untuk bilangan cacah (bilangan positif).
		int, tipe data untuk bilangan bulat (bilangan negatif dan positif).
	*/
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	/*
		Tipe Data Numerik Desimal

		float32
		float64
	*/
	var decimalNumber = 2.62

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	/*
		Tipe Data Boolean

		Tipe data bool berisikan hanya 2 variansi nilai, true dan false
	*/
	var exist bool = true
	fmt.Printf("exist? %t \n", exist)

	/*
		Tipe Data String

		Ciri khas dari tipe data string adalah nilainya di apit oleh
		tanda quote atau petik dua (")
	*/
	var message string = "Halo"
	fmt.Printf("message: %s \n", message)

	var message = `Nama saya "John Wick".
	Salam kenal.
	Mari belajar "Golang".`

	fmt.Println(message)
}
