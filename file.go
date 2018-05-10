/*
	Ada beberapa cara yang bisa digunakan untuk operasi file di Golang.
	Pada bab ini kita akan mempelajari teknik yang paling dasar, yaitu
	dengan memanfaatkan os.File
*/
package main

import (
	"fmt"
	"os"
)

var path = "/Users/novalagung/Documents/temp/test.txt"

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func createFile() {
	// deteksi apakah file sudah ada
	var _, err = os.Stat(path)

	// buat file baru jika belum ada
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("==> file berhasil dibuat", path)
}

func main() {

	/*
		Membuat File Baru
	*/
	createFile()

	/*
		Mengedit Isi File

		Untuk mengedit file, yang perlu dilakukan pertama adalah membuka file dengan level
		akses write. Setelah mendapatkan objek file-nya, gunakan method WriteString() untuk
		pengisian data. Terakhir panggil method Sync() untuk menyimpan perubahan.
	*/
	writeFile()

	/*
		Membaca Isi File

		File yang ingin dibaca harus dibuka terlebih dahulu menggunakan fungsi os.OpenFile()
		dengan level akses minimal adalah read. Setelah itu, gunakan method Read() dengan
		parameter adalah variabel, yang dimana hasil proses baca akan disimpan ke variabel
		tersebut.
	*/
	readFile()

	/*
		Menghapus File

		Cara menghapus file sangatlah mudah, cukup panggil fungsi os.Remove(), masukan path
		file yang ingin dihapus sebagai parameter.
	*/
	deleteFile()
}

func writeFile() {
	// buka file dengan level akses READ & WRITE
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// tulis data ke file
	_, err = file.WriteString("halo\n")
	if isError(err) {
		return
	}
	_, err = file.WriteString("mari belajar golang\n")
	if isError(err) {
		return
	}

	// simpan perubahan
	err = file.Sync()
	if isError(err) {
		return
	}

	fmt.Println("==> file berhasil di isi")
}

func readFile() {
	// buka file
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// baca file
	var text = make([]byte, 1024)
	for {
		n, err := file.Read(text)
		if err != io.EOF {
			if isError(err) {
				break
			}
		}
		if n == 0 {
			break
		}
	}
	if isError(err) {
		return
	}

	fmt.Println("==> file berhasil dibaca")
	fmt.Println(string(text))
}

func deleteFile() {
	var err = os.Remove(path)
	if isError(err) {
		return
	}

	fmt.Println("==> file berhasil di delete")
}
