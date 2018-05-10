package main

import (
	"fmt"
	"strings"
)

type student struct {
	name  string
	grade int
}

/*
	Deklarasi method langsung di tentukan siapa pemiliknya
	pemilik di-mention setelah words func

	method sifatnya sama dengan fungsi normal, perbedaan
	hanya saat deklarasi dan pengaksesan
*/
func (s student) sayHello() {
	fmt.Println("halo", s.name)
}

func (s student) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]
}

func (s student) changeName1(name string) {
	fmt.Println("---> on changeName1, name changed to", name)
	s.name = name
}

func (s *student) changeName2(name string) {
	fmt.Println("---> on changeName2, name changed to", name)
	s.name = name
}

func main() {
	var s1 = student{"john wick", 21}
	s1.sayHello()

	var name = s1.getNameAt(2)
	fmt.Println("nama panggilan :", name)

	/*
		method pointer

		method yang variabel objek pemilik method tersebut berupa pointer.
		Kelebihan method jenis ini adalah, ketika kita melakukan manipulasi
		nilai pada property lain yang masih satu struct, nilai pada property
		tersebut akan di rubah pada reference nya

		Setelah eksekusi statement s1.changeName1("jason bourne"), nilai s1.name
		tidak berubah. Sebenarnya nilainya berubah tapi hanya dalam method changeName1()
		saja, nilai pada reference di objek-nya tidak berubah. Karena itulah ketika objek
		di print value dari s1.name tidak berubah
	*/
	s1.changeName1("jason bourne")
	fmt.Println("s1 after changeName1", s1.name)
	// john wick

	s1.changeName2("ethan hunt")
	fmt.Println("s1 after changeName2", s1.name)
	// ethan hunt

	/*
		Keistimewaan lain method pointer adalah, method itu sendiri bisa dipanggil
		dari objek pointer maupun objek biasa
	*/
	// pengaksesan method dari variabel objek biasa
	var s1 = student{"john wick", 21}
	s1.sayHello()

	// pengaksesan method dari variabel objek pointer
	var s2 = &student{"ethan hunt", 22}
	s2.sayHello()

	/*
		Penggunaan Fungsi strings.Split()
	*/
	strings.Split("ethan hunt", " ")
	// ["ethan", "hunt"]

}
