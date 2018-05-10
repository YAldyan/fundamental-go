package main

import (
	"fmt"
	"math"
)

func main() {
	var diameter float64 = 15
	var area, circumference = calculate(diameter)

	fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
	fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)
}

/*
	Tulis Nilai yang akan dikembalikan float64 float64

	return 2 nilai yang akan dikembalikan
*/

func calculate(d float64) (float64, float64) {
	// hitung luas
	var area = math.Pi * math.Pow(d/2, 2)
	// hitung keliling
	var circumference = math.Pi * d

	// kembalikan 2 nilai
	return area, circumference
}

/*
	Fungsi Dengan Predefined Return Value

	variabel yang digunakan sebagai nilai balik bisa didefinisikan di-awal
	sehingga return tidak perlu disebutkan lagi variabelnya, hanya cukup
	return saja
*/
func calculate(d float64) (area float64, circumference float64) {

	// math.Pow fungis pangkat dari library Math
	area = math.Pi * math.Pow(d/2, 2)

	// math.Pi konstanta built in dari Library Math
	circumference = math.Pi * d

	return
}
