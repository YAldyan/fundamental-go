package main

import (
	"fmt"
)

func main() {
	/*
		Konstanta adalah jenis variabel yang nilainya tidak bisa diubah.
		Inisialisasi nilai hanya dilakukan sekali di awal, setelahnya
		tidak bisa diubah nilainya
	*/

	const firstName string = "Yuhariz Aldyan"
	fmt.Print("Halo ", firstName, "!\n")

	// type inference pada konstanta
	const lastName = "Yuhariz Aldyan"
	fmt.Print("Nice to Meet You ", lastName, "!\n")
}
