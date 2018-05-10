/*
	Goroutine mirip dengan thread thread, tapi sebenarnya bukan.
	Sebuah native thread bisa berisikan sangat banyak goroutine.
	Mungkin lebih pas kalau goroutine disebut sebagai mini thread.
	Goroutine sangat ringan, hanya dibutuhkan sekitar 2kB memori
	saja untuk satu buah goroutine. Ekseksui goroutine bersifat
	asynchronous, menjadikannya tidak saling tunggu dengan goroutine
	lain

	Goroutine merupakan salah satu bagian paling penting dalam Concurrent
	Programming di Golang. Salah satu yang membuat goroutine sangat istimewa
	adalah eksekusi-nya dijalankan di multi core processor. Kita bisa tentukan
	berapa banyak core yang aktif, makin banyak akan makin cepat

	Concurrency atau konkurensi berbeda dengan paralel. Paralel adalah
	eksekusi banyak proses secara bersamaan. Sedangkan konkurensi adalah
	komposisi dari sebuah proses. Konkurensi merupakan struktur, sedangkan
	paralel adalah bagaimana eksekusinya berlangsung.
*/

package main

import (
	"fmt"
	"runtime"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func main() {

	/*
		Fungsi ini digunakan untuk menentukan jumlah core atau processor
		yang digunakan dalam eksekusi program. Jumlah yang diinputkan secara
		otomatis akan disesuaikan dengan jumlah asli logical processor yang
		ada. Jika jumlahnya lebih, maka dianggap menggunakan sejumlah prosesor
		yang ada
	*/
	runtime.GOMAXPROCS(2)

	go print(5, "halo")
	print(5, "apa kabar")

	var input string
	fmt.Scanln(&input)
}
