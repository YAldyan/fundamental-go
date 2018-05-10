/*
	Golang memiliki package encoding/base64, berisikan fungsi-fungsi untuk
	kebutuhan encode dan decode data ke base64 dan sebaliknya. Data yang
	akan di-encode harus bertipe []byte, perlu dilakukan casting untuk data
	-data yang belum sesuai tipenya
*/

package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	/*
		Penerapan Fungsi EncodeToString() & DecodeString()
	*/
	encodeDecodeToString()

	/*
		Penerapan Fungsi Encode() & Decode()
	*/
	encodeDecode()

	/*
		Encode & Decode Data URL
	*/
	encodeDecodeURL()
}

func encodeDecodeToString() {
	var data = "john wick"

	var encodedString = base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("encoded:", encodedString)

	var decodedByte, _ = base64.StdEncoding.DecodeString(encodedString)
	var decodedString = string(decodedByte)
	fmt.Println("decoded:", decodedString)
}

func encodeDecode() {
	var data = "john wick"

	var encoded = make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(encoded, []byte(data))
	var encodedString = string(encoded)
	fmt.Println(encodedString)

	var decoded = make([]byte, base64.StdEncoding.DecodedLen(len(encoded)))
	var _, err = base64.StdEncoding.Decode(decoded, encoded)
	if err != nil {
		fmt.Println(err.Error())
	}
	var decodedString = string(decoded)
	fmt.Println(decodedString)
}

func encodeDecodeURL() {
	var data = "https://kalipare.com/"

	var encodedString = base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(encodedString)

	var decodedByte, _ = base64.URLEncoding.DecodeString(encodedString)
	var decodedString = string(decodedByte)
	fmt.Println(decodedString)
}
