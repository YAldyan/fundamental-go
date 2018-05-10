/*
	Interface kosong atau interface{} adalah tipe data yang sangat spesial.
	Variabel bertipe ini bisa menampung segala jenis data, bahkan array,
	bisa pointer bisa tidak (konsep ini disebut dengan dynamic typing).
*/

package main

import (
	"fmt"
	"strings"
)

type person struct {
	name string
	age  int
}

func main() {

	/*
		Penerapan Interface Kosong sebagai tipe data
	*/
	var secret interface{}

	secret = "ethan hunt"
	fmt.Println(secret)

	secret = []string{"apple", "manggo", "banana"}
	fmt.Println(secret)

	secret = 12.4
	fmt.Println(secret)

	/*
		Membuat Map, key adalah String, sedangkan
		nilainya adalah objeck dengan tipe interface
		kosong
	*/
	var data map[string]interface{}

	data = map[string]interface{}{
		"name":      "ethan hunt",
		"grade":     2,
		"breakfast": []string{"apple", "manggo", "banana"},
	}

	fmt.Println(data)

	/*
		Casting Variabel Interface Kosong
	*/
	var secretX interface{}

	secretX = 2
	var number = secretX.(int) * 10
	fmt.Println(secretX, "multiplied by 10 is :", number)

	secretX = []string{"apple", "manggo", "banana"}
	var gruits = strings.Join(secretX.([]string), ", ")
	fmt.Println(gruits, "is my favorite fruits")

	/*
		Casting Variabel Interface Kosong Ke Objek Pointer
	*/
	var secretY interface{} = &person{name: "wick", age: 27}
	var name = secretY.(*person).name
	fmt.Println(name)

	/*
		Kombinasi Slice, map, dan interface{}
	*/
	var personX = []map[string]interface{}{
		{"name": "Wick", "age": 23},
		{"name": "Ethan", "age": 23},
		{"name": "Bourne", "age": 22},
	}

	for _, each := range personX {
		fmt.Println(each["name"], "age is", each["age"])
	}

	/*
		Slice dengan Interface Kosong
	*/
	var fruits = []interface{}{
		map[string]interface{}{"name": "strawberry", "total": 10},
		[]string{"manggo", "pineapple", "papaya"},
		"orange",
	}

	for _, each := range fruits {
		fmt.Println(each)
	}
}
