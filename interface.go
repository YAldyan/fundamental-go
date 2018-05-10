/*
	Interface adalah kumpulan definisi method yang tidak memiliki isi (hanya definisi saja),
	dan dibungkus dengan nama tertentu.

	Interface merupakan tipe data. Nilai objek bertipe interface default-nya adalah nil.
	Interface mulai bisa digunakan jika sudah ada isinya, yaitu objek konkret yang memiliki
	definisi method minimal sama dengan yang ada di interface-nya.
*/

package main

import (
	"fmt"
	"math"
)

type hitung interface {
	luas() float64
	keliling() float64
}

/*
	Lingkaran
*/

type lingkaran struct {
	diameter float64
}

func (l lingkaran) jariJari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

/*
	Persegi
*/
type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

/*
	embedded interface
*/
type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}

type hitungX interface {
	hitung2d
	hitung3d
}

/*
	Struct Kubus
*/
type kubus struct {
	sisi float64
}

func (k *kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}

func (k *kubus) luas() float64 {
	return math.Pow(k.sisi, 2) * 6
}

func (k *kubus) keliling() float64 {
	return k.sisi * 12
}

func main() {
	var bangunDatar hitung

	bangunDatar = persegi{10.0}
	fmt.Println("===== persegi")
	fmt.Println("luas      :", bangunDatar.luas())
	fmt.Println("keliling  :", bangunDatar.keliling())

	bangunDatar = lingkaran{14.0}
	fmt.Println("===== lingkaran")
	fmt.Println("luas      :", bangunDatar.luas())
	fmt.Println("keliling  :", bangunDatar.keliling())

	/*
		method jarijari tidak ada di interface hitung
		sehingga mesti d casting ke tipe aslinya ling-
		karan sebagaimana teknis penulisannya seperti
		di bawah ini
	*/
	fmt.Println("jari-jari :", bangunDatar.(lingkaran).jariJari())

	/*
		embedded interface

		interface menjadi tipe data dari method di interface lain
	*/
	var bangunRuang hitungX = &kubus{4}

	fmt.Println("===== kubus")
	fmt.Println("luas      :", bangunRuang.luas())
	fmt.Println("keliling  :", bangunRuang.keliling())
	fmt.Println("volume    :", bangunRuang.volume())
}
