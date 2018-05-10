package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {

	/*
		Menampilkan Text di Browser
	*/
	// textInBrowser()

	/*
		Penggunaan Template Web
	*/
	parseTemplate()
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "apa kabar!")
}

func textInBrowser() {

	/*
		Fungsi http.HandleFunc() digunakan untuk routing aplikasi web.
		Maksud dari routing adalah penentuan aksi ketika url tertentu
		diakses oleh user.

		Fungsi http.HandleFunc() memiliki 2 buah parameter yang harus
		diisi. Parameter pertama adalah rute yang diinginkan. Parameter
		kedua adalah callback atau aksi ketika rute tersebut diakses.
		Callback tersebut bertipe fungsi func(w http.ResponseWriter,
		r *http.Request)
	*/
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "halo!")
	})

	/*
		Pada pendaftaran rute /index, callback-nya adalah fungsi index(),
		hal seperti ini diperbolehkan asalkan tipe dari fungsi tersebut
		sesuai.
	*/
	http.HandleFunc("/index", index)

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

func parseTemplate() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]string{
			"Name":    "john wick",
			"Message": "have a nice day",
		}

		var t, err = template.ParseFiles("template.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// tampilkan hasil parsing ke browser
		t.Execute(w, data)
	})

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
