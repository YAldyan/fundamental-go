/*
	Channel digunakan untuk menghubungkan gorutine satu dengan goroutine lain.
	Mekanismenya adalah serah-terima data lewat channel tersebut. Goroutine
	pengirim dan penerima harus berada pada channel yang berbeda (konsep ini
	disebut buffered channel). Pengiriman dan penerimaan data pada channel
	bersifat blocking atau synchronous
*/

package main

import (
	"fmt"
	"runtime"
)

func main() {

	/*
		Penerapan Channel

		Channel merupakan sebuah variabel, dibuat dengan menggunakan keyword make
		dan chan. Variabel channel memiliki tugas menjadi pengirim dan penerima
		data
	*/

	// inisialisasi channel
	var messages = make(chan string)

	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("hello %s", who)
		messages <- data
	}

	go sayHelloTo("john wick")
	go sayHelloTo("ethan hunt")
	go sayHelloTo("jason bourne")

	var message1 = <-messages
	fmt.Println(message1)

	var message2 = <-messages
	fmt.Println(message2)

	var message3 = <-messages
	fmt.Println(message3)

	/*
		Channel Sebagai Tipe Data Parameter
	*/
	runtime.GOMAXPROCS(2)

	var messagesX = make(chan string)

	for _, each := range []string{"wick", "hunt", "bourne"} {
		/*
			goroutine untuk fungsi yang langsung d jalankan.
		*/
		go func(who string) {
			var data = fmt.Sprintf("hello %s", who)
			messagesX <- data
		}(each)
	}

	for i := 0; i < 3; i++ {
		printMessage(messagesX)
	}
}

func printMessage(what chan string) {
	fmt.Println(<-what)
}
