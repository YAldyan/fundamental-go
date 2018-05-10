package main

/*
	web dan JSON, untuk membuat sebuah web API dengan tipe data reponse berbentuk JSON.

	Web API adalah sebuah web yang menerima request dari client dan menghasilkan response,
	biasa berupa JSON/XML. Di bagian ini kita akan buat tanpa authentikasi.
*/

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type student struct {
	ID    string
	Name  string
	Grade int
}

var data = []student{
	student{"E001", "ethan", 21},
	student{"W001", "wick", 22},
	student{"B001", "bourne", 23},
	student{"B002", "bond", 23},
}

/*
	fungsi users() untuk handle rute /users. Didalam fungsi tersebut ada proses deteksi
	jenis request lewat property r.Method(), untuk mencari tahu apakah jenis request
	adalah POST atau GET atau lainnya.
*/
func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var id = r.FormValue("id")
		var result []byte
		var err error

		for _, each := range data {
			if each.ID == id {
				result, err = json.Marshal(each)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.Write(result)
				return
			}
		}

		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}
}

func main() {
	/*
		Pembuatan Web API

		Testing API dengan menjadikan POSTMAN Google Chrome sebagai Client
	*/
	pembuatanWebAPI()
}

func pembuatanWebAPI() {
	http.HandleFunc("/users", users)
	http.HandleFunc("/user", user)

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
