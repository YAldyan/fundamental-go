package main

import (
	"fmt"
)

func main() {

	/*
		Seleksi Kondisi Keyword if, else if & else
	*/

	var pointA = 8

	if pointA == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if pointA > 5 {
		fmt.Println("lulus")
	} else if pointA == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", pointA)
	}

	/*
		Variabel Temporary Pada if - else
	*/
	var pointB = 8840.0

	if percent := pointB / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	/*
		Seleksi Kondisi Menggunakan Keyword switch - case
	*/
	var pointC = 6

	switch pointC {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	/*
		Pemanfaatan case Untuk Banyak Kondisi
	*/
	var pointD = 6

	switch pointD {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	/*
		Switch Dengan Gaya if - else
	*/
	var pointE = 6

	switch {
	case pointE == 8:
		fmt.Println("perfect")
	case (pointE < 8) && (pointE > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	/*
		Penggunaan Keyword fallthrough Dalam switch

		Ketika sebuah case terpenuhi, pengecekkan kondisi
		tidak akan diteruskan ke case-case setelahnya

		Keyword fallthrough digunakan untuk memaksa proses
		pengecekkan diteruskan ke case selanjutnya.
	*/
	var pointF = 6

	switch {
	case pointF == 8:
		fmt.Println("perfect")
	case (pointF < 8) && (pointF > 3):
		fmt.Println("awesome")
		fallthrough
	case pointF < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	/*
		Seleksi Kondisi Bersarang
	*/
	var point = 10

	if point > 7 {
		switch point {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if point == 5 {
			fmt.Println("not bad")
		} else if point == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if point == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}
