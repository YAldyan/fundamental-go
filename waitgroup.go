/*
	Golang menyediakan package sync, berisi cukup banyak API untuk handle masalah
	multiprocessing (goroutine), salah satunya diantaranya adalah yang kita bahas
	di bab ini, yaitu sync.WaitGroup.

	Kegunaan sync.WaitGroup adalah untuk sinkronisasi goroutine. Berbeda dengan
	channel, sync.WaitGroup memang didesain khusus untuk maintain goroutine,
	penggunaannya lebih mudah dan lebih efektif dibanding channel.

	Sebenarnya kurang pas membandingkan sync.WaitGroup dan channel, karena fungsi
	utama dari keduanya adalah berbeda. Channel untuk keperluan sharing data antar
	goroutine, sedangkan sync.WaitGroup untuk sinkronisasi goroutine.
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wait sync.WaitGroup

	// add entities to the WaitGroup
	wait.Add(1)

	go func() {
		fmt.Println("Hello World!")

		// when we use the Done() method, we subtract one
		wait.Done()
	}()

	wait.Wait()

	other_wait()
}

/*
	You might also have noticed that the Goroutines aren't executed in order.
	Of course! We are dealing with a scheduler that doesn't guarantee the order
	of execution of the Goroutines. This is something to keep in mind when
	programming concurrent applications. In fact, if we execute it again, we
	won't necessarily get the same order of output
*/
func other_wait() {
	// variabel untuk waitgroup
	var wait sync.WaitGroup

	goRoutines := 5

	// tambah sejumlah 5 goroutine
	wait.Add(goRoutines)

	for i := 0; i < goRoutines; i++ {
		go func(goRoutineID int) {
			fmt.Printf("ID:%d: Hello goroutines!\n", goRoutineID)
			wait.Done()
		}(i) // parameter goRoutineID
	}

	wait.Wait()
}

/*
	Penerapan sync.WaitGroup

	sync.WaitGroup digunakan untuk menunggu goroutine. Cara penggunaannya sangat mudah, tinggal masukan jumlah goroutine
	yang  dieksekusi, sebagai parameter method Add() object cetakan sync.WaitGroup. Dan pada akhir tiap-tiap goroutine,
	panggil method  one(). Juga, pada baris kode setelah eksekusi goroutine, panggil method Wait().

	Kode di atas merupakan contoh penerapan sync.WaitGroup untuk pengelolahan goroutine. Fungsi doPrint() akan dijalankan
	sebagai goroutine, dengan tugas menampilkan isi variabel message.

	Variabel wg disiapkan bertipe sync.WaitGroup, dibuat untuk sinkronisasi goroutines yang dijalankan.

	Di tiap perulangan statement wg.Add(1) dipanggil. Kode tersebut akan memberikan informasi kepada wg bahwa jumlah
	goroutine  yang sedang diproses ditambah 1 (karena dipanggil 5 kali, maka wg akan sadar bahwa terdapat 5 buah
	goroutine sedang berjalan).

	Di baris selanjutnya, fungsi doPrint() dieksekusi sebagai goroutine. Didalam fungsi tersebut, sebuah method bernama
	Done() dipanggil. Method ini digunakan untuk memberikan informasi kepada wg bahwa goroutine dimana method itu dipanggil
	sudah selesai. Sejumlah 5 buah goroutine dijalankan, maka method tersebut harus dipanggil 5 kali.

	Statement wg.Wait() bersifat blocking, proses eksekusi program tidak akan diteruskan ke baris selanjutnya, sebelum sejumlah
	5 goroutine selesai. Jika Add(1) dipanggil 5 kali, maka Done() juga harus dipanggil 5 kali.

	Beberapa perbedaan antara channel dan sync.WaitGroup:

    1. Channel tergantung kepada goroutine tertentu dalam penggunaannya, tidak seperti sync.WaitGroup yang dia tidak perlu tahu
       goroutine mana saja yang dijalankan, cukup tahu jumlah goroutine yang harus selesai
    2. Penerapan sync.WaitGroup lebih mudah dibanding channel
    3. Kegunaan utama channel adalah untuk komunikasi data antar goroutine. Sifatnya yang blocking bisa kita manfaatkan untuk
       manage goroutine; sedangkan WaitGroup khusus digunakan untuk sinkronisasi goroutine
    4. Performa sync.WaitGroup lebih baik dibanding channel, sumber: https://groups.google.com/forum/#!topic/golang-nuts/
       whpCEk9yLhc

	Kombinasi yang tepat antara sync.WaitGroup dan channel sangat penting, dibutuhkan untuk handle proses concurrent agar bisa
	maksimal.
*/
