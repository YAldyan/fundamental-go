package main

import (
	"fmt"
)

func main() {

	var namesA [4]string
	namesA[0] = "trafalgar"
	namesA[1] = "d"
	namesA[2] = "water"
	namesA[3] = "law"

	fmt.Println(namesA[0], namesA[1], namesA[2], namesA[3])

	/*
		Pengisian Elemen Array yang Melebihi Alokasi Awal
	*/
	var namesB [4]string
	namesB[0] = "trafalgar"
	namesB[1] = "d"
	namesB[2] = "water"
	namesB[3] = "law"
	// namesB[4] = "ez" // baris kode ini menghasilkan error

	/*
		Inisialisasi Nilai Awal Array
	*/
	var fruitsA = [4]string{"apple", "grape", "banana", "melon"}

	fmt.Println("Jumlah element \t\t", len(fruitsA))
	fmt.Println("Isi semua element \t", fruitsA)

	/*
		Inisialisasi Nilai Array Dengan Gaya Vertikal
	*/
	var fruitsB [4]string

	// cara vertikal
	fruitsB = [4]string{
		"apple",
		"grape",
		"banana",
		"melon",
	}

	fmt.Println("Jumlah element \t\t", len(fruitsB))
	fmt.Println("Isi semua element \t", fruitsB)

	/*
		Inisialisasi Nilai Awal Array Tanpa Jumlah Elemen
	*/
	var numbers = [...]int{2, 3, 2, 4, 3}

	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen \t:", len(numbers))

	/*
		Array Multidimensi
	*/
	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	/*
		Perulangan Elemen Array Menggunakan Keyword for
	*/
	var fruitsC = [4]string{"apple", "grape", "banana", "melon"}

	for i := 0; i < len(fruitsC); i++ {
		fmt.Printf("elemen %d : %s\n", i, fruitsC[i])
	}

	/*
		Perulangan Elemen Array Menggunakan Keyword for
	*/
	var fruitsD = [4]string{"apple", "grape", "banana", "melon"}

	for i, fruit := range fruitsD {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}

	/*
		Penggunaan Variabel Underscore _ Dalam for - range

		Kadang kala ketika looping menggunakan for - range,
		ada kemungkinan dimana data yang dibutuhkan adalah
		elemen-nya saja, indeks-nya tidak. Sedangkan kode
		di atas, range mengembalikan 2 data, yaitu indeks
		dan elemen

		Disinilah salah satu kegunaan variabel pengangguran,
		atau underscore (_). Tampung saja nilai yang tidak
		ingin digunakan ke underscore.
	*/
	var fruitsE = [4]string{"apple", "grape", "banana", "melon"}

	for _, fruit := range fruitsE {
		fmt.Printf("nama buah : %s\n", fruit)
	}

}
