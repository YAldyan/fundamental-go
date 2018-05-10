/*
	NoSQL MongoDB

	Golang tidak menyediakan interface generic untuk NoSQL, jadi implementasi driver
	tiap brand NoSQL di Golang bisa berbeda satu dengan lainnya.

	Dari sekian banyak teknologi NoSQL yang ada, yang terpilih untuk dibahas di buku
	ini adalah MongoDB. Dan pada bab ini kita akan belajar cara berkomunikasi dengan
	MongoDB menggunakan driver mgo

	1. Instal mgo menggunakan go get.

	2. Pastikan sudah terinstal MongoDB di komputer anda, dan jangan lupa untuk menjalankan
	   daemon-nya. Jika belum, download dan install terlebih dahulu.

	3. Instal juga MongoDB GUI untuk mempermudah browsing data. Bisa menggunakan MongoChef,
	   Robomongo, atau lainnya.


	Untuk memulai MongoDB
	1. Jalankan Daemon MongoDB di folder bin tempat MongoDB terinstall :

	   mongod.exe --storageEngine=mmapv1 --dbpath "D:\Programming\Database\MongoDB\data"

	2. Jalankan Mongo Client dengan mengetik perintah pada folder tempat MongoDB terinstall : mongo.exe

	3. Membuat database baru. - use [database name] - => use belajar_golang

	4. Membuat Collection baru.
	   mari kita coba buat sebuah koleksi (collection). MongoDB menggunakan istilah koleksi untuk menggantikan tabel.
	   Koleksi berisi kumpulan dokumen atau data dalam format JSON. Kalau di SQL kita menyebutnya dengan record/baris.
       Koleksi bisa dibuat dengan perintah: db.createCollections("nama_koleksi")

*/

package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*
	Tag bson pada property struct dalam konteks mgo, digunakan sebagai penentu nama field
	ketika data disimpan kedalam collection
*/

type student struct {
	Name  string `bson:"name"`
	Grade int    `bson:"Grade"`
}

/*
	Fungsi mgo.Dial() digunakan untuk membuat objek session baru, dengan tipe *mgo.Session.
	Fungsi ini memiliki satu parameter yang harus diisi, yaitu connection string dari server
	mongo yang akan diakses.

	Secara default jenis konsistensi session yang digunakan adalah mgo.Primary . Anda bisa
	mengubahnya lewat method SetMode() milik struct mgo.Session
*/
func connect() (*mgo.Session, error) {
	var session, err = mgo.Dial("localhost")
	if err != nil {
		return nil, err
	}
	return session, nil
}

func insert() {
	var session, err = connect()
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}

	defer session.Close()

	/*
		Struct mgo.Session memiliki method DB() , digunakan untuk memilih database, dan bisa
		langsung di chain dengan fungsi C() untuk memilih collection
	*/
	var collection = session.DB("belajar_golang").C("student")

	/*
		Setelah mendapatkan instance collection-nya, gunakan method Insert() untuk insert data
		ke database. Method ini memiliki parameter variadic, harus diisi pointer data yang ingin
		diinsert.
	*/
	err = collection.Insert(&student{"Wick", 2}, &student{"Ethan", 2})
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}

	fmt.Println("Insert success!")
}

/*
	method Find() milik tipe collection mgo.Collection digunakan untuk melakukan pembacaan data.
	Query selectornya dituliskan menggunakan bson.M lalu disisipkan sebagai parameter fungsi Find() .
	Untuk menggunakan bson.M , package gopkg.in/mgo.v2/bson harus di-import terlebih dahulu.
*/

func find() {
	var session, err = connect()
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}
	defer session.Close()

	var collection = session.DB("belajar_golang").C("student")

	/*
		Variabel result di-inisialisasi menggunakan struct student . Variabel tersebut nantinya
		digunakan untuk menampung hasil pencarian data.
	*/
	var result = student{}

	/*
		query selector ditulis dalam tipe bson.M . Tipe ini sebenarnya adalah alias dari
		map[string]interface{} .
	*/
	var selector = bson.M{"name": "Wick"}

	/*
		Selector tersebut kemudian dimasukan sebagai parameter method Find() , yang kemudian
		di chain langsung dengan method One() untuk mengambil 1 baris datanya. Pointer variabel
		result disisipkan sebagai parameter method tersebut
	*/
	err = collection.Find(selector).One(&result)

	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}

	fmt.Println("Name :", result.Name)
	fmt.Println("Grade :", result.Grade)
}

/*
	Method Update() milik struct mgo.Collection digunakan untuk update data. Ada 2 parameter yang
	harus diisi:
	1. Parameter pertama adalah query selector data yang ingin di update
	2. Parameter kedua adalah data perubahannya

	Di bawah ini adalah contok implementasi method Update() .
*/
func update() {
	var session, err = connect()
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}
	defer session.Close()

	var collection = session.DB("belajar_golang").C("student")

	var selector = bson.M{"name": "Wick"}
	var changes = student{"John Wick", 2}

	err = collection.Update(selector, changes)
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}

	fmt.Println("Update success!")
}

/*
	Cara menghapus document pada collection cukup mudah, tinggal gunakan method Remove()
	dengan isi parameter adalah query selector document yang ingin dihapus
*/
func remove() {
	var session, err = connect()
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}
	defer session.Close()

	var collection = session.DB("belajar_golang").C("student")
	var selector = bson.M{"name": "John Wick"}

	err = collection.Remove(selector)
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}

	fmt.Println("Remove success!")
}

func main() {
	// insert()
	// find()
	// update()
	remove()
}
