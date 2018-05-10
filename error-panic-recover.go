/*
	Error merupakan topik yang penting dalam pemrograman golang.
	Di bagian ini kita akan belajar mengenai pemanfaatan error
	dan cara membuat custom error sendiri. Kita juga akan belajar
	tentang penggunaan panic untuk memunculkan panic error, dan
	recover untuk mengatasinya.
*/

package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	/*
		Pemanfaatan Error.

		error adalah sebuah tipe. Error memiliki memiliki 1 buah
		property berupa method Error(), method ini mengembalikan
		detail pesan error. Error termasuk tipe yang isinya bisa
		kosong atau nilerror adalah sebuah tipe. Error memiliki
		memiliki 1 buah property berupa method Error(), method ini
		mengembalikan detail pesan error. Error termasuk tipe yang
		isinya bisa kosong atau nil
	*/
	var input string
	fmt.Print("Type some number: ")
	fmt.Scanln(&input)

	var number int
	var err error

	/*
		Data pertama (number) berisi hasil konversi. Dan data kedua
		 err, berisi informasi errornya (jika memang terjadi error
		 ketika proses konversi).
	*/
	number, err = strconv.Atoi(input) // konversi inputan menjadi numerik

	if err == nil {
		fmt.Println(number, "is number")
	} else {
		fmt.Println(input, "is not number")
		fmt.Println(err.Error())
	}

	/*
		Membuat Custom Error

		Selain memanfaatkan error hasil kembalian fungsi, kita juga bisa
		membuat error sendiri dengan menggunakan fungsi errors.New (untuk
		menggunakannya harus import package errors terlebih dahulu)
	*/
	var nameX string
	fmt.Print("Type your name: ")
	fmt.Scanln(&nameX)

	if valid, errX := validate(nameX); valid {
		fmt.Println("halo", nameX)
	} else {
		fmt.Println(errX.Error())
	}

	/*
		Penggunaan panic

		Panic digunakan untuk menampilkan trace error sekaligus menghentikan
		flow goroutine (ingat, main juga merupakan goroutine). Setelah ada
		panic, proses selanjutnya tidak di-eksekusi (kecuali proses yang di-defer,
		akan tetap dijalankan tepat sebelum panic muncul).+
	*/
	panic_error()

	/*
		Penggunaan recover

		Recover berguna untuk meng-handle panic error. Pada saat panic muncul,
		recover men-take-over goroutine yang sedang panic (pesan panic tidak
		akan muncul).

		Kita modif sedikit fungsi di-atas untuk mempraktekkan bagaimana cara
		penggunaan recover. Tambahkan fungsi catch(), dalam fungsi ini terdapat
		statement recover() yang dia akan mengembalikan pesan error panic yang
		seharusnya muncul. Fungsi yang berisikan recover() harus di-defer, karena
		meskipun muncul panic, defer tetap di-eksekusi
	*/
	recoverError()

}

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}
	return true, nil
}

func panic_error() {
	var name string
	fmt.Print("Type your name: ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("halo", name)
	} else {
		fmt.Println(err.Error())
		fmt.Println("end")
	}
}

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error occured", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}

func recoverError() {
	defer catch()

	var name string
	fmt.Print("Type your name: ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("halo", name)
	} else {
		panic(err.Error())
		fmt.Println("end")
	}

}
