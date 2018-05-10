package main

import (
	"fmt"
)

func getAverage(numbers []int, ch chan float64) {
	var sum = 0
	for _, e := range numbers {
		sum += e
	}
	ch <- float64(sum) / float64(len(numbers))
}

func getMax(numbers []int, ch chan int) {
	var max = numbers[0]
	for _, e := range numbers {
		if max < e {
			max = e
		}
	}
	ch <- max
}

func main() {
	/*
		Penerapan Keyword Select

		Fungsi utama channel bukan untuk mengontrol goroutine,
		melainkan untuk sharing data antar goroutine. Namun
		channel memang bisa digunakan untuk mengontrol goroutine
	*/

	/*
		Program pencarian rata-rata dan nilai tertinggi berikut
		merupakan contoh sederhana penerapan select dalam channel.
		Akan ada 2 buah goroutine yang masing-masing di-handle oleh
		sebuah channel. Setiap kali goroutine selesai dieksekusi,
		akan dikirimkan datanya ke channel yang bersangkutan. Lalu
		dengan menggunakan select, akan diatur penerimaan datanya.
	*/
	runtime.GOMAXPROCS(2)

	var numbers = []int{3, 4, 3, 5, 6, 3, 2, 2, 6, 3, 4, 6, 3}
	fmt.Println("numbers :", numbers)

	var ch1 = make(chan float64)
	go getAverage(numbers, ch1)

	var ch2 = make(chan int)
	go getMax(numbers, ch2)

	for i := 0; i < 2; i++ {
		select {
		case avg := <-ch1:
			fmt.Printf("Avg \t: %.2f \n", avg)
		case max := <-ch2:
			fmt.Printf("Max \t: %d \n", max)
		}
	}
}
