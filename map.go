package main

import (
	"fmt"
)

func main() {

	/*
		Map

		Tipe Key-Value
	*/

	/*
		Cara Penggunaan
	*/

	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"]) // januari 50
	fmt.Println("mei", chicken["mei"])         // mei 0

	/*
		Inisialisasi Nilai Map
	*/
	// cara vertikal
	var chicken1 = map[string]int{"januari": 50, "februari": 40}

	// cara horizontal
	var chicken2 = map[string]int{
		"januari":  50,
		"februari": 40,
	}

	/*
		Inisialisasi Tanpa Parameter

		Khusus inisialisasi data menggunakan keyword new, yang dihasilkan
		adalah data pointer. Untuk mengambil nilai aslinya bisa dengan
		menggunakan tanda asterisk (*)
	*/
	var chicken3 = map[string]int{}
	var chicken4 = make(map[string]int)
	var chicken5 = *new(map[string]int)

	/*
		Iterasi Item Map Menggunakan for - range
	*/
	var chicken = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"april":    67,
	}

	for key, val := range chicken {
		fmt.Println(key, "  \t:", val)
	}

	/*
		Menghapus Item Map

		Fungsi delete() digunakan untuk menghapus item dengan key
		tertentu pada variabel map
	*/
	var chicken = map[string]int{"januari": 50, "februari": 40}

	fmt.Println(len(chicken)) // 2
	fmt.Println(chicken)

	delete(chicken, "januari")

	fmt.Println(len(chicken)) // 1
	fmt.Println(chicken)

	/*
		Deteksi Keberadaan Item Dengan Key Tertentu
	*/

	var chicken = map[string]int{"januari": 50, "februari": 40}
	var value, isExist = chicken["mei"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}

	/*
		Kombinasi Slice & Map

		Cara menggunakannya cukup mudah, contohnya seperti
		[]map[string]int, artinya slice yang tipe tiap elemen
		-nya adalah map[string]int.
	*/

	var chickens = []map[string]string{
		map[string]string{"name": "chicken blue", "gender": "male"},
		map[string]string{"name": "chicken red", "gender": "male"},
		map[string]string{"name": "chicken yellow", "gender": "female"},
	}

	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}

}
