package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

var baseURL = "http://localhost:8080"

type student struct {
	ID    string
	Name  string
	Grade int
}

func fetchUsers() ([]student, error) {
	var err error
	var client = &http.Client{}
	var data []student

	/*
		Fungsi http.NewRequest() digunakan untuk membuat request baru.
		Fungsi tersebut memiliki 3 parameter yang wajib diisi.

		Parameter pertama, berisikan tipe request POST atau GET atau lainnya
		Parameter kedua, adalah URL tujuan request
		Parameter ketiga, form data request (jika ada)
	*/
	request, err := http.NewRequest("POST", baseURL+"/users", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	/*
		data response perlu di-close setelah tidak dipakai. Caranya seperti pada
		kode defer response.Body.Close().
	*/
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func fetchUser(ID string) (student, error) {
	var err error
	var client = &http.Client{}
	var data student

	/*
		Statement url.Values{} akan menghasilkan objek yang nantinya digunakan
		sebagai form data request. Pada objek tersebut perlu di set data apa
		saja yang ingin dikirimkan menggunakan fungsi Set() seperti pada param.
		Set("id", ID).
	*/
	var param = url.Values{}
	param.Set("id", ID)

	/*
		Statement bytes.NewBufferString(param.Encode()) maksudnya, objek form
		data di-encode lalu diubah menjadi bentuk bytes.Buffer, yang nantinya
		disisipkan pada parameter ketiga pemanggilan fungsi http.NewRequest()
	*/
	var payload = bytes.NewBufferString(param.Encode())

	request, err := http.NewRequest("POST", baseURL+"/user", payload)
	if err != nil {
		return data, err
	}

	/*
		Karena data yang akan dikirim di-encode, maka pada header perlu di set tipe
		konten request-nya. Kode request.Header.Set("Content-Type", "application/x-
		www-form-urlencoded") artinya tipe konten request di set sebagai application
		/x-www-form-urlencoded.
	*/
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		return data, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func main() {
	/*
		Penggunaan HTTP Request Tanpa Form Data
	*/
	httpRequestTanpaFormData()

	/*
		Penggunaan HTTP Request Dengan Form Data
	*/
	httpRequestDenganFormData()
}

func httpRequestTanpaFormData() {
	var users, err = fetchUsers()
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}

	for _, each := range users {
		fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", each.ID, each.Name, each.Grade)
	}
}

/*
	Response dari rute /user bukan berupa array objek, melainkan sebuah objek. Maka pada
	saat decode pastikan tipe variabel penampung hasil decode data response adalah student
	(bukan []student).
*/
func httpRequestDenganFormData() {
	var user1, err = fetchUser("W001")
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}

	fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", user1.ID, user1.Name, user1.Grade)
}
