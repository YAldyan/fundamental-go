/*
	Hash adalah algoritma enkripsi untuk mengubah text menjadi deretan karakter acak.
	Jumlah karakter hasil hash selalu sama. Hash termasuk one-way encription, membuat
	hasil dari hash tidak bisa dikembalikan ke text asli.

	SHA1 atau Secure Hash Algorithm 1 merupakan salah satu algoritma hashing yang sering
	digunakan untuk enkripsi data. Hasil dari sha1 adalah data dengan lebar 20 byte atau
	160 bit, biasa ditampilkan dalam bentuk bilangan heksadesimal 40 digit.
*/
package main

import (
	"crypto/sha1"
	"fmt"
	"time"
)

func main() {

	//  Golang menyediakan package crypto/sha1, berisikan library untuk keperluan hashing

	/*
		Penerapan Hash SHA1
	*/
	penerapanHash()

	/*
		Metode Salting Pada Hash SHA1
	*/
	var text = "this is secret"
	fmt.Printf("original : %s\n\n", text)

	var hashed1, salt1 = doHashUsingSalt(text)
	fmt.Printf("hashed 1 : %s\n\n", hashed1)
	// 929fd8b1e58afca1ebbe30beac3b84e63882ee1a

	var hashed2, salt2 = doHashUsingSalt(text)
	fmt.Printf("hashed 2 : %s\n\n", hashed2)
	// cda603d95286f0aece4b3e1749abe7128a4eed78

	var hashed3, salt3 = doHashUsingSalt(text)
	fmt.Printf("hashed 3 : %s\n\n", hashed3)
	// 9e2b514bca911cb76f7630da50a99d4f4bb200b4

	_, _, _ = salt1, salt2, salt3

}

/*
	Variabel hasil dari sha1.New() adalah objek bertipe hash.Hash, memiliki dua buah method
	Write() dan Sum().

	Method Write() digunakan untuk menge-set data yang akan di-hash. Data harus dalam bentuk
	[]byte.

	Method Sum() digunakan untuk eksekusi proses hash, menghasilkan data yang sudah di-hash
	dalam bentuk []byte. Method ini membutuhkan sebuah parameter, isi dengan nil.

	Untuk mengambil bentuk heksadesimal string dari data yang sudah di-hash, bisa memanfaatkan
	fungsi fmt.Sprintf dengan layout format %x
*/
func penerapanHash() {
	var text = "this is secret"
	var sha = sha1.New()
	sha.Write([]byte(text))
	var encrypted = sha.Sum(nil)
	var encryptedString = fmt.Sprintf("%x", encrypted)

	fmt.Println(encryptedString)
	// f4ebfd7a42d9a43a536e2bed9ee4974abf8f8dc8
}

/*
	Salt dalam konteks kriptografi adalah data acak yang digabungkan pada data asli sebelum
	proses hash dilakukan.

	Hash merupakan enkripsi satu arah dengan lebar data yang sudah pasti, sangat mungkin
	sekali kalau hasil hash untuk beberapa data adalah sama. Disinilah kegunaan salt, teknik
	ini berguna untuk mencegah serangan menggunakan metode pencocokan data-data yang hasil
	hash-nya adalah sama (dictionary attack).
*/
func doHashUsingSalt(text string) (string, string) {
	var salt = fmt.Sprintf("%d", time.Now().UnixNano())
	var saltedText = fmt.Sprintf("text: '%s', salt: %s", text, salt)
	fmt.Println(saltedText)
	var sha = sha1.New()
	sha.Write([]byte(saltedText))
	var encrypted = sha.Sum(nil)

	return fmt.Sprintf("%x", encrypted), salt
}
