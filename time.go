/*
	Time disini maksudnya adalah gabungan date dan time, bukan hanya waktu saja
*/

package main

import (
	"fmt"
	"time"
)

func main() {

	/*
		Penggunaan time.Time
	*/
	// ambil waktu sekarang
	var time1 = time.Now()
	fmt.Printf("time1 %v\n", time1)
	// time1 2015-09-01 17:59:31.73600891 +0700 WIB

	/*
		waktu ditentukan sendiri dengan parameter
		time.Date(tahun, bulan, tanggal, jam, menit, detik, nanodetik, timezone)
	*/
	var time2 = time.Date(2011, 12, 24, 10, 20, 0, 0, time.UTC)
	fmt.Printf("time2 %v\n", time2)
	// time2 2011-12-24 10:20:00 +0000 UTC

	/*
		Method Milik time.Time
	*/
	var now = time.Now()
	fmt.Println("year:", now.Year(), "month:", now.Month())
	// year: 2015 month: 8
	/*
		now.Year()			int				Tahun
		now.YearDay()		int				Hari ke-? di mulai awal tahun
		now.Month()			int				Bulan
		now.Weekday()		string			Nama hari. Bisa menggunakan now.Weekday().String()
											untuk mengambil bentuk string-nya
		now.ISOWeek()		(int, int)		Tahun dan minggu ke-? mulai awal tahun
		now.Day()			int				Tanggal
		now.Hour()			int				Jam
		now.Minute()		int				Menit
		now.Second()		int				Detik
		now.Nanosecond()	int				Nano detik
		now.Local()			time.Time		Waktu dalam timezone lokal
		now.Location()		*time.Location	Mengambil informasi lokasi, apakah local atau utc.
											Bisa menggunakan now.Location().String() untuk
											mengambil bentuk string-nya
		now.Zone()			(string, int)	Mengembalikan informasi timezone offset dalam string
											dan numerik. Sebagai contoh WIB, 25200
		now.IsZero()		bool			Deteksi apakah nilai object now adalah 01 Januari
											tahun 1, 00:00:00 UTC. Jika iya maka bernilai true
		now.UTC()			time.Time		Waktu dalam timezone UTC
		now.Unix()			int64			Waktu dalam format unix time
		now.UnixNano()		int64			Waktu dalam format unix time. Infomasi nano detik juga
											imasukkan
		now.String()		string			Waktu dalam string
	*/

	/*
		Parsing time.Time
	*/
	var layoutFormat, value string
	var date time.Time

	layoutFormat = "2006-01-02 15:04:05"
	value = "2015-09-02 08:04:00"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t->", date.String())
	// 2015-09-02 08:04:00 +0000 UTC

	layoutFormat = "02/01/2006 MST"
	value = "02/09/2015 WIB"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t\t->", date.String())
	// 2015-09-02 00:00:00 +0700 WIB

	/*
		2006			Tahun 4 digit								2015
		006				Tahun 3 digit								015
		06				Tahun 2 digit								15
		01				Bulan 2 digit								05
		1				Bulan 1 digit jika dibawah bulan 10, 		5, 12
						selainnya 2 digit
		January			Nama bulan dalam bahasa inggris				September, August
		Jan				Nama bulan dalam bahasa inggris, 3 huruf	Sep, Aug
		02				Tanggal 2 digit								02
		2				Tanggal 1 digit jika dibawah bulan 10, 		8, 31
						selainnya 2 digit
		Monday			Nama hari dalam bahasa inggris				Saturday, Friday
		Mon				Nama hari dalam bahasa inggris, 3 huruf		Sat, Fri
		15				Jam dengan format 24 jam					18
		03				Jam dengan format 12 jam 2 digit			05, 11
		3				Jam dengan format 12 jam 1 digit jika
						dibawah jam 11, selainnya 2 digit			5, 11
		PM				AM/PM, biasa digunakan dengan format
						jam 12 jam									PM, AM
		04				Menit 2 digit								08
		4				Menit 1 digit jika dibawah menit 10,
						selainnya 2 digit							8, 24
		05				Detik 2 digit								06
		5				Detik 1 digit jika dibawah detik 10,
						selainnya 2 digit							6, 36
		999999			Nano detik									124006
		MST				Lokasi timezone								UTC, WIB, EST
		Z0700			Offset timezone								Z, +0700, -0200

	*/

	/*
		Predefined Layout Format Untuk Keperluan Parsing Time

		Salah satu predefined layout yang bisa digunakan adalah time.RFC822.
		Format ini sama dengan 02 Jan 06 15:04 MST
	*/
	var dateX, _ = time.Parse(time.RFC822, "02 Sep 15 08:00 WIB")
	fmt.Println(dateX.String())
	// 2015-09-02 08:00:00 +0700 WIB

	/*
		Predefined Layout Format			Layout Format
		time.ANSIC							Mon Jan _2 15:04:05 2006
		time.UnixDate						Mon Jan _2 15:04:05 MST 2006
		time.RubyDate						Mon Jan 02 15:04:05 -0700 2006
		time.RFC822							02 Jan 06 15:04 MST
		time.RFC822Z						02 Jan 06 15:04 -0700
		time.RFC850							Monday, 02-Jan-06 15:04:05 MST
		time.RFC1123						Mon, 02 Jan 2006 15:04:05 MST
		time.RFC1123Z						Mon, 02 Jan 2006 15:04:05 -0700
		time.RFC3339						2006-01-02T15:04:05Z07:00
		time.RFC3339Nano					2006-01-02T15:04:05.999999999Z07:00
		time.Kitchen						3:04PM
		time.Stamp							Jan _2 15:04:05
		time.StampMilli						Jan _2 15:04:05.000
		time.StampMicro						Jan _2 15:04:05.000000
		time.StampNano						Jan _2 15:04:05.000000000
	*/

	/*
		Format time.Time

		Method Format() milik tipe time.Time digunakan untuk membentuk
		output string sesuai dengan layout format yang diinginkan
	*/
	var dateY, _ = time.Parse(time.RFC822, "02 Sep 15 08:00 WIB")

	var dateS1 = dateY.Format("Monday 02, January 2006 15:04 MST")
	fmt.Println("dateS1", dateS1)
	// Wednesday 02, September 2015 08:00 WIB

	var dateS2 = dateY.Format(time.RFC3339)
	fmt.Println("dateS2", dateS2)
	// 2015-09-02T08:00:00+07:00

	/*
		Handle Error Parsing time.Time
	*/
	var dateZ, err = time.Parse("06 Jan 15", "02 Sep 15 08:00 WIB")

	if err != nil {
		fmt.Println("error", err.Error())
		return
	}

	fmt.Println(dateZ)
}
