package main

import "testing"

var (
	kubus              Kubus   = Kubus{4}
	volumeSeharusnya   float64 = 64
	luasSeharusnya     float64 = 96
	kelilingSeharusnya float64 = 48
)

func TestHitungVolume(t *testing.T) {
	t.Logf("Volume : %.2f", kubus.Volume())

	if kubus.Volume() != volumeSeharusnya {
		t.Errorf("SALAH! harusnya %.2f", volumeSeharusnya)
	}
}

func TestHitungLuas(t *testing.T) {
	t.Logf("Luas : %.2f", kubus.Luas())

	if kubus.Luas() != luasSeharusnya {
		t.Errorf("SALAH! harusnya %.2f", luasSeharusnya)
	}
}

func TestHitungKeliling(t *testing.T) {
	t.Logf("Keliling : %.2f", kubus.Keliling())

	if kubus.Keliling() != kelilingSeharusnya {
		t.Errorf("SALAH! harusnya %.2f", kelilingSeharusnya)
	}
}

/*
	command

	testing		: go test unit-test.go go test unit-test.go
	benchmark	:
*/

/*
	Method Test

	Method			Kegunaan
	Log()			Menampilkan log
	Logf()			Menampilkan log menggunakan format
	Fail()			Menandakan terjadi Fail() dan proses testing fungsi tetap diteruskan
	FailNow()		Menandakan terjadi Fail() dan proses testing fungsi dihentikan
	Failed()		Menampilkan laporan fail
	Error()			Log() diikuti dengan Fail()
	Errorf()		Logf() diikuti dengan Fail()
	Fatal()			Log() diikuti dengan failNow()
	Fatalf()		Logf() diikuti dengan failNow()
	Skip()			Log() diikuti dengan SkipNow()
	Skipf()			Logf() diikuti dengan SkipNow()
	SkipNow()		Menghentikan proses testing fungsi, dilanjutkan ke testing fungsi
					setelahnya
	Skiped()		Menampilkan laporan skip
	Parallel()		Menge-set bahwa eksekusi testing adalah parallel

*/
