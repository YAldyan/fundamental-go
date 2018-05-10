package main

import (
	"fmt"
	"strconv"
)

func main() {

	/*
		Konversi Menggunakan strconv
	*/

	/*
		Fungsi strconv.Atoi()
	*/
	stringToInt()

	/*
		Fungsi strconv.Itoa()
	*/
	intToString()

	/*
		Fungsi strconv.ParseInt()
	*/
	stringDesimalToNumeric()
	stringBinerToNumeric()

	/*
		Fungsi strconv.FormatInt()

		Berguna untuk konversi data numerik int64 ke string dengan
		basis numerik bisa ditentukan sendiri
	*/
	stringFormatInt()

	/*
		Fungsi strconv.ParseFloat()

		untuk konversi string ke numerik desimal dengan lebar data
		bisa ditentukan

		string "24.12" dikonversi ke float dengan lebar tipe data
		float32.
	*/
	stringToParseFloat()

	/*
		Fungsi strconv.FormatFloat()

		untuk konversi data bertipe float64 ke string dengan format
		eksponen, lebar digit desimal, dan lebar tipe data bisa ditentukan.
	*/
	stringFormatFloat()

	/*
		Fungsi strconv.FormatBool()

		untuk konversi bool ke string
	*/
	boolToString()

	/*
		Konversi Data Menggunakan Casting
	*/
	var a float64 = float64(24)
	fmt.Println(a) // 24

	var b int32 = int32(24.00)
	fmt.Println(b) // 24

	/*
		Casting string ↔ byte

		kode ascii masing-masing karakter string
		dibaca sebagai byte-nya
	*/
	stringToByte()
	byteToString()

	/*
		Konversi Data interface{} Menggunakan Teknik Type Assertions

		Type assertions merupakan teknik casting data interface{} ke
		segala jenis tipe (dengan syarat data tersebut memang bisa di-
		casting ke tipe tujuan).
	*/
	var data = map[string]interface{}{
		"nama":    "john wick",
		"grade":   2,
		"height":  156.5,
		"isMale":  true,
		"hobbies": []string{"eating", "sleeping"},
	}

	fmt.Println(data["nama"].(string))
	fmt.Println(data["grade"].(int))
	fmt.Println(data["height"].(float64))
	fmt.Println(data["isMale"].(bool))
	fmt.Println(data["hobbies"].([]string))

	/*
		Tipe asli data pada variabel interface{} bisa diketahui dengan
		cara meng-casting ke tipe type

		casting ke tipe type hanya bisa dilakukan pada switch
	*/
	for _, val := range data {
		switch val.(type) {
		case string:
			fmt.Println(val.(string))
		case int:
			fmt.Println(val.(int))
		case float64:
			fmt.Println(val.(float64))
		case bool:
			fmt.Println(val.(bool))
		case []string:
			fmt.Println(val.([]string))
		default:
			fmt.Println(val.(int))
		}
	}
}

func stringToInt() {
	var str = "124"
	var num, err = strconv.Atoi(str)

	if err == nil {
		fmt.Println(num) // 124
	}
}

func intToString() {
	var num = 124
	var str = strconv.Itoa(num)

	fmt.Println(str) // "124"
}

func stringDesimalToNumeric() {
	var str = "124"
	/*
		str		: string yang akan dikonversi ke numeric
		10		: basis nilai (desimal)
		64		: lebar nilai int64
	*/
	var num, err = strconv.ParseInt(str, 10, 64)

	if err == nil {
		fmt.Println(num) // 124
	}
}

func stringBinerToNumeric() {
	var str = "1010"
	/*
		str		: string yang akan dikonversi ke numeric
		2		: basis nilai (biner)
		8		: lebar nilai int8
	*/
	var num, err = strconv.ParseInt(str, 2, 8)

	if err == nil {
		fmt.Println(num) // 10
	}
}

func stringFormatInt() {
	var num = int64(24)
	var str = strconv.FormatInt(num, 8)

	fmt.Println(str) // 30
}

func stringToParseFloat() {
	var str = "24.12"
	var num, err = strconv.ParseFloat(str, 32)

	if err == nil {
		fmt.Println(num) // 24.1200008392334
	}
}

func stringFormatFloat() {
	var num = float64(24.12)
	var str = strconv.FormatFloat(num, 'f', 6, 64)

	fmt.Println(str) // 24.120000
}

func boolToString() {
	var bul = true
	var str = strconv.FormatBool(bul)

	fmt.Println(str) // 124
}

func stringToByte() {
	var text1 = "halo"
	var b = []byte(text1)

	fmt.Printf("%d %d %d %d \n", b[0], b[1], b[2], b[3])
	// 104 97 108 111
}

func byteToString() {
	var byte1 = []byte{104, 97, 108, 111}
	var s = string(byte1)

	fmt.Printf("%s \n", s)
	// halo
}
