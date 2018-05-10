package main

import (
	"fmt"
)

/*
	Definisi Closure adalah sebuah fungsi yang bisa disimpan dalam variabel.
	Dengan menerapkan konsep tersebut, kita bisa membuat fungsi didalam fungsi,
	atau bahkan membuat fungsi yang mengembalikan fungsi. Closure merupakan anonymous
	function atau fungsi tanpa nama. Biasa dimanfaatkan untuk membungkus suatu proses
	yang hanya dipakai sekali atau dipakai pada blok tertentu saja.
*/

func main() {

	/*
		Closure Disimpan Sebagai Variabel
	*/
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	var numbers = []int{2, 3, 4, 3, 4, 2, 3}
	var min, max = getMinMax(numbers)

	/*
		Template %v digunakan untuk menampilkan segala jenis data.
		Bisa array, int, float, bool, dan lainnya
	*/
	fmt.Printf("data : %v\nmin  : %v\nmax  : %v\n", numbers, min, max)

	/*
		Immediately-Invoked Function Expression (IIFE)
	*/
	var numbers2 = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	var newNumbers = func(min int) []int {
		var r []int
		for _, e := range numbers2 {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return r
	}(3)

	fmt.Println("original number :", numbers2)
	fmt.Println("filtered number :", newNumbers)

	/*
		Closure Sebagai Nilai Kembalian
	*/
	var max2 = 3
	var numbers3 = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	var howMany, getNumbers = findMax(numbers3, max2)
	var theNumbers = getNumbers()

	fmt.Println("numbers\t:", numbers3)
	fmt.Printf("find \t: %d\n\n", max2)

	fmt.Println("found \t:", howMany)    // 9
	fmt.Println("value \t:", theNumbers) // [2 3 0 3 2 0 2 0 3]
}

func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}
	return len(res), func() []int {
		return res
	}
}
