/*
	Penerimaan data lewat channel yang jumlah goroutine-nya banyak bisa
	lebih mudah dengan memanfaatkan keyword for - range

	for - range jika diterapkan pada channel berfungsi untuk handle
	penerimaan data. Setiap kali ada pengiriman data ke channel,
	perulangan akan berjalan sekali, setelah 1 perulangan selesai
	akan menunggu lagi penerimaan data selanjutnya. Perulangan hanya
	akan berhenti jika channel di close atau di-non-aktifkan. Fungsi
	close digunakan utuk me-non-aktifkan channel

	Channel yang sudah di-close tidak bisa digunakan lagi untuk menerima
	maupun mengirim data, itulah kenapa perulangan pada for - range juga
	ikut berhenti.
*/

/*

 */
package main

import (
	"fmt"
	"runtime"
)

func sendMessage(ch chan<- string) {
	for i := 0; i < 20; i++ {
		ch <- fmt.Sprintf("data %d", i)
	}
	close(ch)
}

func printMessage(ch <-chan string) {
	for message := range ch {
		fmt.Println(message)
	}
}

func main() {
	/*
		Penerapan for - range - close Pada Channel

		for - range untuk penerimaan data dari channel
	*/
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)
	go sendMessage(messages)
	printMessage(messages)

	/*
		Channel Direction

		Ada yang unik dengan fitur parameter channel yang disediakan Golang.
		Level akses channel bisa ditentukan, apakah hanya sebagai penerima,
		pengirim, atau penerima sekaligus pengirim. Konsep ini disebut dengan
		channel direction

		ch chan string		Parameter ch bisa digunakan untuk mengirim dan menerima data
		ch chan<- string	Parameter ch hanya bisa digunakan untuk mengirim data
		ch <-chan string	Parameter ch hanya bisa digunakan untuk menerima data

		Sebagai contoh fungsi sendMessage(ch chan<- string) yang parameter ch
		dideklarasikan dengan level akses untuk pengiriman data saja. Channel
		tersebut hanya bisa digunakan untuk mengirim,
		contohnya: ch <- fmt.Sprintf("data %d", i).

		Dan sebaliknya pada fungsi printMessage(ch <-chan string), channel ch hanya
		bisa digunakan untuk menerima data saja.
	*/
}
