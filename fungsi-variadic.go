/*
	Golang mengadopsi konsep variadic function atau pembuatan fungsi
	dengan parameter sejenis yang tak terbatas. Maksud tak terbatas
	disini adalah jumlah parameter yang disisipkan ketika pemanggilan
	fungsi bisa berapa saja.
*/

package main

import (
	"fmt"
	"strings"
)

func main() {

	/*
		penerapan fungsi variadic
	*/
	var avg = calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)

	// nilai dikembalikan dalam bentuk string
	var msg = fmt.Sprintf("Rata-rata : %.2f", avg)
	fmt.Println(msg)

	/*
		Pengisian Parameter Fungsi Variadic Menggunakan Data Slice
	*/
	var numbers = []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}
	var avg2 = calculate(numbers...)
	var msg2 = fmt.Sprintf("Rata-rata : %.2f", avg2)

	fmt.Println(msg2)

	/*
		Fungsi Dengan Parameter Biasa & Variadic
	*/
	var hobbies = []string{"sleeping", "eating"}
	yourHobbies("wick", hobbies...)
}

/*
	parameter tidak terbatas
	nilai diakses dengan index
	mirip array dan slice
*/
func calculate(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	/*
		operasi bilangan, perkalian, pembagian
		penambahan, dan pengurangan hanya bisa
		dilakukan dengan tipe data sejenis
	*/
	var avg = float64(total) / float64(len(numbers))
	return avg
}

func yourHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")

	fmt.Printf("Hello, my name is: %s\n", name)
	fmt.Printf("My hobbies are: %s\n", hobbiesAsString)
}
