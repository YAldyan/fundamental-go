package main

import (
	"fmt"
)

func main() {
	/*
		slice adalah reference(alamat) elemen
		array

		slice bisa diperoleh dari manipulasi
		array/slice lainnya.

		artinya, ketika kita copy slice, yg
		di copy adalah alamat elemet array
		atau slice lainnya, sehingga nilai
		atau element yang didapati adalah
		sama, karena yg ditunjuk oleh slice
		baru adalah alamat slice/array lama
		yang dimanipulasi.
	*/

	/*
		inisialisasi slice

		mengkapling alamat memori dan diisi
		dengan nilai/elemen array.
	*/
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits[0]) // "apple"

	/*
		Slice Merupakan Tipe Data Reference

		ketika slice baru diubah nilainya, slice
		lama yang menjadi sumber copian akan ikut
		berubah karena yang ditunjuk adalah alamat
		memori.
	*/
	var fruits = []string{"apple", "grape", "banana", "melon"}

	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]

	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]

	fmt.Println(fruits)   // [apple grape banana melon]
	fmt.Println(aFruits)  // [apple grape banana]
	fmt.Println(bFruits)  // [grape banana melon]
	fmt.Println(aaFruits) // [grape]
	fmt.Println(baFruits) // [grape]

	// Buah "grape" diubah menjadi "pinnaple"
	baFruits[0] = "pinnaple"

	fmt.Println(fruits)   // [apple pinnaple banana melon]
	fmt.Println(aFruits)  // [apple pinnaple banana]
	fmt.Println(bFruits)  // [pinnaple banana melon]
	fmt.Println(aaFruits) // [pinnaple]
	fmt.Println(baFruits) // [pinnaple]

	/*
		fungsi len()
	*/
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(len(fruits)) // 4

	/*
		fungsi cap()

		Fungsi cap() digunakan untuk menghitung lebar atau kapasitas
		maksimum slice. Nilai kembalian fungsi ini untuk slice yang
		baru dibuat pasti sama dengan len, tapi bisa berubah seiring
		operasi slice yang dilakukan
	*/
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(len(fruits)) // len: 4
	fmt.Println(cap(fruits)) // cap: 4

	/*
		Slicing yang dimulai dari indeks 0 hingga 3 akan mengembalikan
		elemen-elemen mulai indeks 0 hingga sebelum indeks 3, dengan
		lebar kapasitas adalah sama dengan slice aslinya
	*/
	var aFruits = fruits[0:3]
	fmt.Println(len(aFruits)) // len: 3
	fmt.Println(cap(aFruits)) // cap: 4

	/*
		fruits(x,y)

		Sedangkan slicing yang dimulai dari indeks x, yang dimana nilai
		x adalah lebih dari 0, membuat elemen ke-x slice yang diambil
		menjadi elemen ke-0 slice baru. Hal inilah yang membuat kapasitas
		slice berubah.
	*/
	var bFruits = fruits[1:4]
	fmt.Println(len(bFruits)) // len: 3
	fmt.Println(cap(bFruits)) // cap: 3

	/*
		Fungsi append()

		Fungsi append() digunakan untuk menambahkan elemen pada slice.
		Elemen baru tersebut diposisikan setelah indeks paling akhir.
		Nilai balik fungsi ini adalah slice yang sudah ditambahkan nilai
		barunya
	*/
	var fruits = []string{"apple", "grape", "banana"}
	var cFruits = append(fruits, "papaya")

	fmt.Println(fruits)  // ["apple", "grape", "banana"]
	fmt.Println(cFruits) // ["apple", "grape", "banana", "papaya"]

	/*
		keterangan.

		Ketika jumlah elemen dan lebar kapasitas adalah sama
		(len(fruits) == cap(fruits)), maka elemen baru hasil append()
		merupakan referensi baru.

		Ketika jumlah elemen lebih kecil dibanding kapasitas
		(len(fruits) < cap(fruits)), elemen baru tersebut ditempatkan
		kedalam cakupan kapasitas, menjadikan semua elemen slice lain
		yang referensi-nya sama akan berubah nilainya.
	*/
	var fruits = []string{"apple", "grape", "banana"}
	var bFruits = fruits[0:2]

	fmt.Println(cap(bFruits)) // 3
	fmt.Println(len(bFruits)) // 2

	fmt.Println(fruits)  // ["apple", "grape", "banana"]
	fmt.Println(bFruits) // ["apple", "grape"]

	var cFruits = append(bFruits, "papaya")

	fmt.Println(fruits)  // ["apple", "grape", "papaya"]
	fmt.Println(bFruits) // ["apple", "grape"]
	fmt.Println(cFruits) // ["apple", "grape", "papaya"]

	/*
		Fungsi copy()

		Fungsi copy() digunakan untuk men-copy elemen slice
		pada parameter ke-2, untuk digabungkan dengan slice
		pada parameter ke-1
	*/

	var fruits = []string{"apple"}
	var aFruits = []string{"watermelon", "pinnaple"}

	var copiedElemen = copy(fruits, aFruits)

	fmt.Println(fruits)       // ["apple", "watermelon", "pinnaple"]
	fmt.Println(aFruits)      // ["watermelon", "pinnaple"]
	fmt.Println(copiedElemen) // 1

	/*
		Pengaksesan Elemen Slice Dengan 3 Indeks

		indeks ketiga merupakan maksimum capacity
		daripada slice yang akan sedang dimanipulasi
	*/

	var fruits = []string{"apple", "grape", "banana"}
	var aFruits = fruits[0:2]
	var bFruits = fruits[0:2:2]

	fmt.Println(fruits)      // ["apple", "grape", "banana"]
	fmt.Println(len(fruits)) // len: 3
	fmt.Println(cap(fruits)) // cap: 3

	fmt.Println(aFruits)      // ["apple", "grape"]
	fmt.Println(len(aFruits)) // len: 2
	fmt.Println(cap(aFruits)) // cap: 3

	fmt.Println(bFruits)      // ["apple", "grape"]
	fmt.Println(len(bFruits)) // len: 2
	fmt.Println(cap(bFruits)) // cap: 2
}
