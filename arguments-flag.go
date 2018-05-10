/*
	Arguments adalah data opsional yang disisipkan ketika eksekusi program.
	Sedangkan flag merupakan ekstensi dari argument. Dengan flag, penulisan
	argument menjadi lebih rapi dan terstruktur.
*/
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	/*
		Penggunaan Arguments

		Pada saat eksekusi program disisipkan juga argument-nya. Sebagai contoh disisipkan
		3 buah data sebagai argumen, yaitu: banana, potato, dan ice cream.

		Untuk eksekusinya sendiri bisa menggunakan go run.

		Menggunakan go run
			$ go run argument-flag.go banana potato "ice cream"
	
		Variabel os.Args mengembalikan tak hanya arguments saja, tapi juga path file executable
		(jika eksekusi-nya menggunakan go run maka path akan merujuk ke folder temporary). Gunakan
		os.Args[1:] untuk mengambil slice arguments-nya saja.
	*/
	penggunaanArguments()

	/*
		Penggunaan Flag

		Flag memiliki kegunaan yang sama seperti arguments, yaitu untuk parameterize eksekusi
		program, dengan penulisan dalam bentuk key-value

		go run argument-flag.go.go -name="john wick" -age=28

		Nama Fungsi										Return Value

		flag.Bool(name, defaultValue, usage)			*bool
		flag.Duration(name, defaultValue, usage)		*time.Duration
		flag.Float64(name, defaultValue, usage)			*float64
		flag.Int(name, defaultValue, usage)				*int
		flag.Int64(name, defaultValue, usage)			*int64
		flag.String(name, defaultValue, usage)			*string
		flag.Uint(name, defaultValue, usage)			*uint
		flag.Uint64(name, defaultValue, usage)			*uint64
	*/
	penggunaanFlag()

	/*
		Deklarasi Flag Dengan Cara Passing Reference Variabel Penampung Data

		Sebenarnya ada 2 cara deklarasi flag yang bisa digunakan, dan cara di atas merupakan
		cara pertama.

		Cara kedua mirip dengan cara pertama, perbedannya adalah kalau di cara pertama nilai
		pointer flag dikembalikan lalu ditampung variabel; sedangkan pada cara kedua, nilainya
		diambil lewat parameter pointer.
	*/
	referencenFlag()
}

func penggunaanArguments() {

	/*
		Data arguments bisa didapat lewat variabel os.Args (package os perlu
		di-import terlebih dahulu)
	*/
	var argsRaw = os.Args
	fmt.Printf("-> %#v\n", argsRaw)

	var args = argsRaw[1:]
	fmt.Printf("-> %#v\n", args)
	// []string{"banana", "potato"}
}

// go run argument-flag.go -name="john wick" -age=28
func penggunaanFlag() {
	/*
		Kode tersebut maksudnya adalah, disiapkan flag bertipe string, dengan key
		adalah name, dengan nilai default "anonymous", dan keterangan "type your name"

		Nilai balik fungsi flag.String() adalah string pointer, jadi perlu di-dereference
		terlebih dahulu agar bisa mendapatkan nilai aslinya (*dataName).
	*/
	var name = flag.String("name", "anonymous", "type your name")
	var age = flag.Int64("age", 25, "type your age")

	flag.Parse()
	fmt.Printf("name\t: %s\n", *name)
	fmt.Printf("age\t: %d\n", *age)
}

/*
	Tinggal tambahkan suffix Var pada pemanggilan nama fungsi flag yang digunakan (contoh
	flag.IntVar(), flag.BoolVar(), dll), lalu disisipkan referensi variabel penampung flag
	sebagai parameter pertama.

	Kegunaan dari parameter terakhir method-method flag adalah untuk memunculkan hints atau
	petunjuk arguments apa saja yang bisa dipakai, ketika argument --help ditambahkan saat
	eksekusi program.
*/
func referencenFlag() {
	// cara ke-1
	var data1 = flag.String("name", "anonymous", "type your name")
	fmt.Println(*data1)

	// cara ke-2
	var data2 string
	flag.StringVar(&data2, "gender", "male", "type your gender")
	fmt.Println(data2)
}
