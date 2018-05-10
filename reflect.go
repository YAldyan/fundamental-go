/*
	Reflection adalah teknik untuk inspeksi variabel, mengambil informasi dari variabel
	tersebut atau bahkan memanipulasinya. Cakupan informasi yang bisa didapatkan lewat
	reflection sangat luas, seperti melihat struktur variabel, tipe, nilai pointer, dan
	banyak lagi.
*/

package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Grade int
}

func (s *student) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(s)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("nama      :", reflectType.Field(i).Name)
		fmt.Println("tipe data :", reflectType.Field(i).Type)
		fmt.Println("nilai     :", reflectValue.Field(i).Interface())
		fmt.Println("")
	}
}

func (s *student) SetName(name string) {
	s.Name = name
}

func main() {
	/*
		Mencari Tipe Data & Value Menggunakan Reflect
	*/
	var number = 23
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("tipe  variabel :", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variabel :", reflectValue.Int())
	}

	/*
		Pengaksesan Nilai Dalam Bentuk interface{}
	*/
	var numberX = 23
	var reflectValueX = reflect.ValueOf(numberX)

	fmt.Println("tipe  variabel :", reflectValueX.Type())

	/*
		Fungsi Interface() mengembalikan nilai interface kosong atau interface{}.
	*/
	fmt.Println("nilai variabel :", reflectValueX.Interface())

	/*
		Nilai aslinya sendiri bisa diakses dengan meng-casting interface kosong tersebut
	*/
	var nilai = reflectValue.Interface().(int)
	fmt.Println("Nilai Asli : ", nilai)

	/*
		Pengaksesan Informasi Property Variabel Objek
	*/
	var s1 = &student{Name: "wick", Grade: 2}
	s1.getPropertyInfo()

	/*
		Pengaksesan Informasi Method Variabel Objek

		memanggil method object dengan reflection
	*/
	var sY = &student{Name: "john wick", Grade: 2}
	fmt.Println("nama :", sY.Name)

	var reflectValueY = reflect.ValueOf(sY)
	var method = reflectValueY.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("wick"),
	})

	fmt.Println("nama :", sY.Name)

}
