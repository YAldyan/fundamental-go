package main

import (
	"fmt"
)

/*
	deklarasi struct
*/
type student struct {
	name  string
	grade int
}

type studentX struct {
	person
	grade int
}

type person struct {
	name string
	age  int
}

type personY struct {
	name string
	age  int
}

type studentY struct {
	person
	age   int
	grade int
}

/*
	Nested struct

	Nested struct adalah anonymous struct yang di-embed
	ke sebuah struct. Deklarasinya langsung didalam struct
	peng-embed

	Teknik ini biasa digunakan ketika decoding data json
	yang struktur datanya cukup kompleks dengan proses
	decode hanya sekali.
*/
type studentG struct {
	person struct {
		name string
		age  int
	}
	grade   int
	hobbies []string
}

/*
	Tag property dalam struct

	Tag biasa dimanfaatkan untuk keperluan encode/decode
	data json. Informasi tag juga bisa diakses lewat reflect
*/
type personG struct {
	name string "tag1"
	age  int    "tag2"
}

func main() {
	/*
		struct ini adalah mirip dengan class
		kalo di pemrograman high level lain.
	*/

	/*
		penerapan struct
	*/
	var s1 student
	s1.name = "john wick"
	s1.grade = 2

	fmt.Println("name  :", s1.name)
	fmt.Println("grade :", s1.grade)

	/*
		Inisialisasi Object Struct
	*/

	var s2 = student{"ethan", 2}

	/*
		Pada deklarasi s3, dilakukan juga pengisian property
		ketika pencetakan objek. Hanya saja, yang diisi hanya
		name saja. Cara ini cukup efektif jika digunakan untuk
		membuat objek baru yang nilai property-nya tidak semua
		harus disiapkan di awal. Keistimewaan lain menggunakan
		cara ini adalah penentuan nilai property bisa dilakukan
		dengan tidak berurutan

		var s4 = student{name: "wayne", grade: 2}
		var s5 = student{grade: 2, name: "bruce"}
	*/
	var s3 = student{name: "jason"}

	fmt.Println("student 2 :", s2.name)
	fmt.Println("student 3 :", s3.name)

	/*
		Variabel Objek Pointer

		objeck tipe pointer bisa menyimpan nilai struct
	*/
	var s4 = student{name: "wick", grade: 2}

	var s5 *student = &s1
	fmt.Println("student 4, name :", s4.name)
	fmt.Println("student 5, name :", s5.name)

	s5.name = "ethan"
	fmt.Println("student 4, name :", s4.name)
	fmt.Println("student 5, name :", s5.name)

	/*
		Embedded Struct

		objek struct menjadi property dari Struct Lainnya
	*/

	var sX = studentX{}
	sX.name = "wick"
	sX.age = 21
	sX.grade = 2

	fmt.Println("name  :", sX.name)
	fmt.Println("age   :", sX.age)
	fmt.Println("age   :", sX.person.age)
	fmt.Println("grade :", sX.grade)

	/*
		Embedded Struct

		dengan nama property yg sama

		mengakses age milik personY dan studentY
		harus dilakukan secara eksplisit karena
		keduanya memiliki property dengan nama
		yang sama
	*/
	var sY = studentY{}
	sY.name = "wick"
	sY.age = 21        // age of student
	sY.person.age = 22 // age of person

	fmt.Println(sY.name)
	fmt.Println(sY.age)
	fmt.Println(sY.person.age)

	/*
		Pengisian Nilai Sub-Struct
	*/
	var pZ = person{name: "wick", age: 21}
	var sZ = studentX{person: pZ, grade: 2}

	fmt.Println("name  :", sZ.name)
	fmt.Println("age   :", sZ.age)
	fmt.Println("grade :", sZ.grade)

	/*
		Anonymous Struct

		Anonymous struct adalah struct yang tidak dideklarasikan di awal,
		melainkan ketika dibutuhkan saja, langsung pada saat penciptaan objek.
		Teknik ini cukup efisien untuk pembuatan variabel objek yang struct nya
		hanya dipakai sekali
	*/

	// anonymous struct tanpa pengisian property
	var sA = struct {
		person
		grade int
	}{}

	// anonymous struct dengan pengisian property
	var sB = struct {
		person
		grade int
	}{
		person: person{"wick", 21},
		grade:  2,
	}

	sA.person = person{"wick", 21}
	sA.grade = 2

	fmt.Println("name  :", sA.person.name)
	fmt.Println("age   :", sA.person.age)
	fmt.Println("grade :", sA.grade)

	fmt.Println("name  :", sB.person.name)
	fmt.Println("age   :", sB.person.age)
	fmt.Println("grade :", sB.grade)

	/*
		Kombinasi Slice & Struct
	*/
	var allStudents = []person{
		{name: "Wick", age: 23},
		{name: "Ethan", age: 23},
		{name: "Bourne", age: 22},
	}

	for _, student := range allStudents {
		fmt.Println(student.name, "age is", student.age)
	}

	/*
		Inisialisasi Slice Anonymous Struct

		Anonymous struct bisa dijadikan sebagai tipe sebuah slice.
		Dan nilai awalnya juga bisa diinisialisasi langsung pada
		saat deklarasi
	*/
	var allStudentS = []struct {
		person
		grade int
	}{
		{person: person{"wick", 21}, grade: 2},
		{person: person{"ethan", 22}, grade: 3},
		{person: person{"bond", 21}, grade: 3},
	}

	for _, student := range allStudentS {
		fmt.Println(student)
	}

	/*
		Deklarasi Anonymous Struct Menggunakan Keyword var
	*/
	var studentH struct {
		person
		grade int
	}

	studentH.person = person{"wick", 21}
	studentH.grade = 2
}
