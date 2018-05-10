/*
	Regex atau regular expression adalah suatu teknik yang digunakan untuk
	pencocokan string dengan pola tertentu. Regex biasa dimanfaatkan untuk
	pencarian dan pengubahan data string
*/

package main

import (
	"fmt"
	"regexp"
)

func main() {

	/*
		Penerapan Regexp
	*/
	penerapanRegex()

	/*
		Method MatchString()

		untuk mendeteksi apakah string memenuhi sebuah pola regexp
	*/
	matchString()

	/*
		Method FindString()

		untuk mencari string yang memenuhi kriteria regexp yang telah
		ditentukan
	*/
	findString()

	/*
		Method FindStringIndex()

		untuk mencari index string kembalian hasil dari operasi regexp

		FindString() hanya saja yang dikembalikan indeks-nya.
	*/
	findStringIndex()

	/*
		Method FindAllString()

		untuk mencari banyak string yang memenuhi kriteria regexp yang
		telah ditentukan.

		Jumlah data yang dikembalikan bisa ditentukan. Jika diisi dengan
		-1, maka akan mengembalikan semua data.
	*/
	findAllString()

	/*
		Method ReplaceAllString()

		untuk me-replace semua string yang memenuhi kriteri regexp, dengan
		string lain.+
	*/
	replaceAllString()

	/*
		Method ReplaceAllStringFunc()

		untuk me-replace semua string yang memenuhi kriteri regexp, dengan
		kondisi yang bisa ditentukan untuk setiap substring yang akan di
		replace.
	*/
	replaceAllStringFunc()
}

func penerapanRegex() {

	/*
		Ekspresi [a-z]+ maknanya adalah, semua string yang merupakan alphabet yang hurufnya
		kecil. Ekspresi tersebut di-compile oleh regexp.Compile() lalu disimpan ke variabel
		objek regex bertipe regexp.*Regexp.
	*/
	var text = "banana burger soup"
	var regex, err = regexp.Compile(`[a-z]+`)

	if err != nil {
		fmt.Println(err.Error())
	}

	/*
		Struct regexp.Regexp memiliki banyak method, salah satunya adalah FindAllString(),
		berfungsi untuk mencari semua string yang sesuai dengan ekspresi regex, dengan
		kembalian berupa array string.

		Jumlah hasil pencarian dari regex.FindAllString() bisa ditentukan. Contohnya pada
		res1, ditentukan maksimal 2 data saja pada nilai kembalian. Jika batas di set -1,
		maka akan mengembalikan semua data.
	*/

	var res1 = regex.FindAllString(text, 2)
	fmt.Printf("%#v \n", res1)
	// ["banana", "burger"]

	var res2 = regex.FindAllString(text, -1)
	fmt.Printf("%#v \n", res2)
	// ["banana", "burger", "soup"]
}

func matchString() {
	var text = "banana burger soup"
	var regex, _ = regexp.Compile(`[a-z]+`)

	var isMatch = regex.MatchString(text)
	fmt.Println(isMatch)
	// true
}

func findString() {
	var text = "banana burger soup"
	var regex, _ = regexp.Compile(`[a-z]+`)

	var str = regex.FindString(text)
	fmt.Println(str)
	// "banana"
}

func findStringIndex() {
	var text = "banana burger soup"
	var regex, _ = regexp.Compile(`[a-z]+`)

	var idx = regex.FindStringIndex(text)
	fmt.Println(idx)
	// [0, 6]

	var str = text[0:6]
	fmt.Println(str)
	// "banana"
}

func findAllString() {
	var text = "banana burger soup"
	var regex, _ = regexp.Compile(`[a-z]+`)

	var str1 = regex.FindAllString(text, -1)
	fmt.Println(str1)
	// ["banana", "burger", "soup"]

	var str2 = regex.FindAllString(text, 1)
	fmt.Println(str2)
	// ["banana"]
}

func replaceAllString() {
	var text = "banana burger soup"
	var regex, _ = regexp.Compile(`[a-z]+`)

	var str = regex.ReplaceAllString(text, "potato")
	fmt.Println(str)
	// "potato potato potato"
}

/*
	Pada contoh di atas, jika salah satu substring yang match adalah "burger" maka akan
	diganti dengan "potato", string selainnya tidak di replace.
*/
func replaceAllStringFunc() {
	var text = "banana burger soup"
	var regex, _ = regexp.Compile(`[a-z]+`)

	var str = regex.ReplaceAllStringFunc(text, func(each string) string {
		if each == "burger" {
			return "potato"
		}
		return each
	})
	fmt.Println(str)
	// "banana potato soup"
}
