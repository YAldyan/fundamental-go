/*
	Golang menyediakan package strings, berisikan cukup banyak
	fungsi untuk keperluan pengolahan data string
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	//Fungsi strings.Contains()
	var isExists = strings.Contains("john wick", "wick")
	fmt.Println(isExists)

	//strings.HasPrefix()
	var isPrefix1 = strings.HasPrefix("john wick", "jo")
	fmt.Println(isPrefix1) // true

	var isPrefix2 = strings.HasPrefix("john wick", "wi")
	fmt.Println(isPrefix2) // false

	//strings.HasSuffix()
	var isSuffix1 = strings.HasSuffix("john wick", "ic")
	fmt.Println(isSuffix1) // false

	var isSuffix2 = strings.HasSuffix("john wick", "ck")
	fmt.Println(isSuffix2) // true

	//Fungsi strings.Count()
	var howMany = strings.Count("ethan hunt", "t")
	fmt.Println(howMany) // 2

	//Fungsi strings.Index()
	var index1 = strings.Index("ethan hunt", "ha")
	fmt.Println(index1) // 2
	var index2 = strings.Index("ethan hunt", "n")
	fmt.Println(index2) // 4

	//Fungsi strings.Replace()
	var text = "banana"
	var find = "a"
	var replaceWith = "o"

	var newText1 = strings.Replace(text, find, replaceWith, 1)
	fmt.Println(newText1) // "bonana"

	var newText2 = strings.Replace(text, find, replaceWith, 2)
	fmt.Println(newText2) // "bonona"

	var newText3 = strings.Replace(text, find, replaceWith, -1)
	fmt.Println(newText3) // "bonono"

	//Fungsi strings.Repeat()
	var str = strings.Repeat("na", 4)
	fmt.Println(str) // "nananana"

	//Fungsi strings.Split()
	var string1 = strings.Split("the dark knight", " ")
	fmt.Println(string1) // ["the", "dark", "knight"]

	var string2 = strings.Split("batman", "")
	fmt.Println(string2) // ["b", "a", "t", "m", "a", "n"]

	//Fungsi strings.Join()
	var data = []string{"banana", "papaya", "tomato"}
	var str = strings.Join(data, "-")
	fmt.Println(str) // "banana-papaya-tomato"

	//Fungsi strings.ToLower()
	var str = strings.ToLower("aLAy")
	fmt.Println(str) // "alay"

	//Fungsi strings.ToUpper()
	var str = strings.ToUpper("eat!")
	fmt.Println(str) // "EAT!"
}
