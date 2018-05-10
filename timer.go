package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	/*
		Fungsi time.Sleep()
	*/
	fmt.Println("start")
	time.Sleep(time.Second * 4)
	fmt.Println("after 4 seconds")

	/*
		Fungsi time.NewTimer()

		Fungsi ini sedikit berbeda dengan time.Sleep().
		Fungsi time.NewTimer() mengembalikan objek bertipe
		*time.Timer yang memiliki property C. Cara kerja
		fungsi ini, setelah jeda waktu yang ditentukan
		sebuah data akan dikirimkan lewat channel C.
		Penggunaan fungsi ini harus diikuti dengan statement
		untuk penerimaan data dari channel C
	*/
	var timer = time.NewTimer(4 * time.Second)
	fmt.Println("start")
	<-timer.C
	fmt.Println("finish")

	/*
		Fungsi time.AfterFunc()

		Fungsi time.AfterFunc() memiliki 2 parameter. Parameter
		pertama adalah durasi timer, dan parameter kedua adalah
		callback nya. Callback tersebut akan dieksekusi jika waktu
		sudah memenuhi durasi timer
	*/
	var ch = make(chan bool)

	time.AfterFunc(4*time.Second, func() {
		fmt.Println("expired")
		ch <- true
	})

	fmt.Println("start")
	<-ch
	fmt.Println("finish")

	/*
		Beberapa hal yang perlu diketahui dalam menggunakan fungsi ini:

		- Jika tidak ada serah terima data lewat channel, maka eksekusi
		time.AfterFunc() adalah asynchronous dan tidak blocking.

		- Jika ada serah terima data lewat channel, maka fungsi akan tetap
		berjalan asynchronous dan tidak blocking hingga baris kode dimana
		penerimaan data channel dilakukan
	*/

	/*
		Fungsi time.After()

		Perbedaannya adalah, fungsi timer.After() akan mengembalikan data
		channel, sehingga perlu menggunakan tanda <- dalam penerapannya
	*/
	<-time.After(4 * time.Second)
	fmt.Println("expired")

	/*
		Kombinasi Timer & Goroutine
	*/
	var timeout = 5
	var chX = make(chan bool)

	go timerX(timeout, chX)
	go watcher(timeout, chX)

	var input string
	fmt.Print("what is 725/25 ? ")
	fmt.Scan(&input)

	if input == "29" {
		fmt.Println("the answer is right!")
	} else {
		fmt.Println("the answer is wrong!")
	}
}

func timerX(timeout int, ch chan<- bool) {
	time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		ch <- true
	})
}

func watcher(timeout int, ch <-chan bool) {
	<-ch
	fmt.Println("\ntime out! no answer more than", timeout, "seconds")
	os.Exit(0)
}
