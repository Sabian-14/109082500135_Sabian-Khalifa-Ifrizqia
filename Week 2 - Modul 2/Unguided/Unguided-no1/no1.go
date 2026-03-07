package main

import "fmt"

func main() {
	var tahun int

	fmt.Print("Masukkan tahun : ")
	fmt.Scan(&tahun)

	kabisat := false

	if (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0) {
		kabisat = true
	}

	fmt.Println("Kabisat :", kabisat)
}