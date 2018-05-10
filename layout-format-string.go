package main

import (
	"fmt"
)

type student struct {
	name        string
	height      float64
	age         int32
	isGraduated bool
	hobbies     []string
}

var data = student{
	name:        "wick",
	height:      182.5,
	age:         26,
	isGraduated: false,
	hobbies:     []string{"eating", "sleeping"},
}

/*
	Layout Format

	https://dasarpemrogramangolang.novalagung.com/37-string-format.html
*/

func main() {
	// Layout Format %b
	fmt.Printf("%b\n", data.age)

	// Layout Format %c
	// Layout Format %d
	// Layout Format %e atau %E
	// Layout Format %f atau %F
	// Layout Format %g atau %G
	// Layout Format %o
	// Layout Format %p
	// Layout Format %q
	// Layout Format %s
	// Layout Format %t
	// Layout Format %T
	// Layout Format %v
	// Layout Format %+v

	// Layout Format %#v
	/*
		Digunakan untuk memformat struct, mengembalikan
		nama dan nilai tiap property sesuai dengan struktur
		truct dan juga bagaimana objek tersebut dideklarasikan
	*/
	fmt.Printf("%#v\n", data)
	/*
		main.student{name:"wick", height:182.5, age:26, isGraduated:false,
		hobbies:[]string{"eating", "sleeping"}}
	*/

	// Layout Format %x atau %X
	/*
		Digunakan untuk memformat data numerik,
		menjadi bentuk string numerik berbasis
		16 (heksadesimal).
	*/
	fmt.Printf("%x\n", data.age)
	// 1a

	/*
		Jika digunakan pada tipe data string,
		maka akan mengembalikan kode heksadesimal
		tiap karakter.
	*/
	var d = data.name

	fmt.Printf("%x%x%x%x\n", d[0], d[1], d[2], d[3])
	// 7769636b

	fmt.Printf("%x\n", d)
	// 7769636b

	// Layout Format %%
	fmt.Printf("%%\n")
	// %
}
