/*
	Defer digunakan untuk mengakhirkan eksekusi sebuah statement.
	Sedangkan Exit digunakan untuk menghentikan program (ingat,
	menghentikan program, tidak seperti return yang menghentikan blok kode).
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		Penerapan keyword defer

		Seperti yang sudah dijelaskan secara singkat di atas,
		bahwa defer digunakan untuk mengakhirkan eksekusi baris kode.
		Ketika eksekusi blok sudah selesai, statement yang di defer baru akan dijalankan

		defer dieksekusi diakhir, ketika sebuah blok kode sdh d eksekusi.
	*/
	defer fmt.Println("halo")
	fmt.Println("selamat datang")

	orderSomeFood("pizza")
	orderSomeFood("burger")

	/*
		Penerapan Fungsi os.Exit()

		Exit digunakan untuk menghentikan program secara paksa pada saat itu juga.
		Semua statement setelah exit tidak akan di eksekusi, termasuk juga defer
	*/
	defer fmt.Println("halo")
	os.Exit(1)
	fmt.Println("selamat datang")
}

/*
	Statement yang di-defer akan tetap muncul meskipun blok
	kode diberhentikan ditengah jalan menggunakan return
*/
func orderSomeFood(menu string) {
	defer fmt.Println("Terimakasih, silakan tunggu")
	if menu == "pizza" {
		fmt.Print("Pilihan tepat!", " ")
		fmt.Print("Pizza ditempat kami paling enak!", "\n")
		return
	}

	fmt.Println("Pesanan anda:", menu)
}
