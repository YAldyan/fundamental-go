package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	id    string
	name  string
	age   int
	grade int
}

func connect() (*sql.DB, error) {

	// user:password@tcp(host:port)/dbname
	// user@tcp(host:port)/dbname
	// user     => root
	// password =>
	// host     => 127.0.0.1 atau localhost
	// port     => 3306
	// dbname   => db_belajar_golang
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_belajar_golang")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func sqlQuery() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var age = 27
	rows, err := db.Query("select id, name, grade from tb_student where age = ?", age)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []student

	for rows.Next() {
		var each = student{}
		var err = rows.Scan(&each.id, &each.name, &each.grade)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each.name)
	}
}

func sqlQueryRow() {
	var db, err = connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var result = student{}
	var id = "E001"
	err = db.
		QueryRow("select name, grade from tb_student where id = ?", id).
		Scan(&result.name, &result.grade)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("\nname: %s\ngrade: %d\n", result.name, result.grade)
}

func sqlPrepare() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	/*
		Method Prepare() digunakan untuk deklarasi query, yang mengembalikan objek bertipe
		sql.*Stmt
	*/
	stmt, err := db.Prepare("select name, grade from tb_student where id = ?")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var result1 = student{}
	stmt.QueryRow("E001").Scan(&result1.name, &result1.grade)
	fmt.Printf("\nname: %s\ngrade: %d\n", result1.name, result1.grade)

	var result2 = student{}
	stmt.QueryRow("W001").Scan(&result2.name, &result2.grade)
	fmt.Printf("\nname: %s\ngrade: %d\n", result2.name, result2.grade)

	var result3 = student{}
	stmt.QueryRow("B001").Scan(&result3.name, &result3.grade)
	fmt.Printf("\nname: %s\ngrade: %d\n", result3.name, result3.grade)
}

func sqlExec() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("insert into tb_student values (?, ?, ?, ?)", "G001", "Galahad", 29, 2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("insert success!")

	_, err = db.Exec("update tb_student set age = ? where id = ?", 28, "G001")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("update success!")

	_, err = db.Exec("delete from tb_student where id = ?", "G001")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("delete success!")
}

func main() {
	/*
		Membaca Data Dari MySQL Server
	*/
	sqlQuery()

	/*
		Membaca 1 Record Data Menggunakan Method QueryRow()s
	*/
	sqlQueryRow()

	/*
		Eksekusi Query Menggunakan Prepare()

		eknik prepared statement adalah teknik penulisan query di awal dengan kelebihan
		bisa di re-use atau digunakan banyak kali untuk eksekusi yang berbeda-beda

		Metode ini bisa digabung dengan Query() maupun QueryRow()
	*/
	sqlPrepare()

	/*
		Insert, Update, & Delete Data Menggunakan Exec()

		Untuk operasi insert, update, dan delete; dianjurkan untuk tidak menggunakan
		fungsi sql.Query() ataupun sql.QueryRow() untuk eksekusinya. Direkomendasikan
		eksekusi perintah-perintah tersebut lewat fungsi Exec()
	*/
	sqlExec()

	/*
		Koneksi Dengan Engine Database Lain

		sql.Open(driverName, connectionString)

		Sebagai contoh saya menggunakan driver pq untuk koneksi ke server Postgres, maka
		connection string-nya:
			sql.Open("pq", "user=postgres password=secret dbname=test sslmode=disable")

		https://github.com/golang/go/wiki/SQLDrivers
	*/

}
