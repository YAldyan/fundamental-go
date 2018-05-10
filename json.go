package main

import (
	"encoding/json"
	"fmt"
)

/*
	Perlu diketahui bahwa untuk decode data json ke variabel objek hasil struct,
	semua level akses property struct-nya harus publik.
*/
type User struct {
	/*
		FullName memiliki tag json:"Name". Tag tersebut digunakan untuk mapping data
		json ke property yang bersangkutan.
	*/
	FullName string `json:"Name"`
	Age      int
}

func main() {
	/*
		Decode JSON Ke Variabel Objek Cetakan Struct

		Di Golang data json tipenya adalah string. Dengan menggunakan json.Unmarshal,
		json string bisa dikonversi menjadi bentuk objek, entah itu dalam bentuk map
		[string]interface{} ataupun variabel objek hasil struct
	*/
	decodeJSONtoStruct()

	/*
		Decode JSON Ke map[string]interface{} & interface{}

		Tak hanya ke objek cetakan struct, target decoding data json juga bisa berupa
		variabel bertipe map[string]interface{}.
	*/
	var jsonString = `{"Name": "john wick", "Age": 27}`
	var jsonData = []byte(jsonString)

	var data1 map[string]interface{}
	json.Unmarshal(jsonData, &data1)

	fmt.Println("user :", data1["Name"])
	fmt.Println("age  :", data1["Age"])

	/*
		Variabel bertipe interface{} juga bisa digunakan untuk menampung hasil decode.
		Dengan catatan pada pengaksesan nilai property, harus dilakukan casting terlebih
		dahulu ke map[string]interface{}.
	*/
	var data2 interface{}
	json.Unmarshal(jsonData, &data2)

	var decodedData = data2.(map[string]interface{})
	fmt.Println("user :", decodedData["Name"])
	fmt.Println("age  :", decodedData["Age"])

	/*
		Decode Array JSON Ke Array Objek

		Decode data dari array json ke slice/array objek masih sama, siapkan saja variabel
		penampung hasil decode dengan tipe slice struct
	*/
	arrayJSONtoArrayObjek()

	/*
		Encode Objek Ke JSON

		Fungsi json.Marshal digunakan untuk decoding data ke json string. Sumber data bisa
		berupa variabel objek cetakan struct, map[string]interface{}, atau slice.
	*/
	encodetoJSONString()
}

func decodeJSONtoStruct() {
	var jsonString = `{"Name": "john wick", "Age": 27}`
	var jsonData = []byte(jsonString)

	var data User

	/*
		Dalam penggunaan fungsi json.Unmarshal, variabel penampung hasil decode harus
		di-pass dalam bentuk pointer, contohnya seperti &data
	*/
	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user :", data.FullName)
	fmt.Println("age  :", data.Age)
}

func arrayJSONtoArrayObjek() {
	var jsonString = `[
    	{"Name": "john wick", "Age": 27},
    	{"Name": "ethan hunt", "Age": 32}
	]`

	var data []User

	var err = json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user 1:", data[0].FullName)
	fmt.Println("user 2:", data[1].FullName)

}

func encodetoJSONString() {
	var object = []User{{"john wick", 27}, {"ethan hunt", 32}}
	var jsonData, err = json.Marshal(object)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var jsonString = string(jsonData)
	fmt.Println(jsonString)
}
