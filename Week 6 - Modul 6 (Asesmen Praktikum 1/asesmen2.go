package main

import "fmt"

func tanggunganHari(jumlahHari int, tujuan string) int {
	if tujuan == "domestik" {
		if jumlahHari > 3 {
			return 3
		}
		return jumlahHari
	} else {
		if jumlahHari > 8 {
			return 8
		}
		return jumlahHari
	}
}

func biayaPerHari(jumlahMhs int) int {
	biayaMakan := 2 * 35000
	biayaPenginapan := 250000
	uangSaku := 300000

	totalPerMhs := biayaMakan + biayaPenginapan + uangSaku
	return totalPerMhs * jumlahMhs
}

func hitungTotalBiaya(jumlahMhs, lamaHari int, tujuan string) int {
	hari := tanggunganHari(lamaHari, tujuan)
	biayaHarian := biayaPerHari(jumlahMhs)

	if tujuan == "mancanegara" {
		biayaHarian = biayaHarian * 3 / 2
	}

	return hari * biayaHarian
}

func main() {
	var jumlah int
	var lama int
	var tujuan string

	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scan(&jumlah)

	fmt.Print("Masukkan lama hari study tour: ")
	fmt.Scan(&lama)

	fmt.Print("Masukkan tujuan study tour (domestik/mancanegara): ")
	fmt.Scan(&tujuan)

	total := hitungTotalBiaya(jumlah, lama, tujuan)

	fmt.Println("Biaya perjalanan yang harus dikeluarkan Tel-U : Rp.", total)
}