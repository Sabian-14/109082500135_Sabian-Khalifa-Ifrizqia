package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type angka int
type kata string

type Buku struct {
	judul        kata
	penulis      kata
	penerbit     kata
	tahunTerbit  angka
	jumlahHalaman angka
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var b Buku

	fmt.Println("=== INPUT BIODATA BUKU ===")

	fmt.Print("Masukkan judul buku : ")
	judul, _ := reader.ReadString('\n')
	b.judul = kata(strings.TrimSpace(judul))

	fmt.Print("Masukkan nama penulis : ")
	penulis, _ := reader.ReadString('\n')
	b.penulis = kata(strings.TrimSpace(penulis))

	fmt.Print("Masukkan penerbit : ")
	penerbit, _ := reader.ReadString('\n')
	b.penerbit = kata(strings.TrimSpace(penerbit))

	fmt.Print("Masukkan tahun terbit : ")
	fmt.Scanln(&b.tahunTerbit)

	fmt.Print("Masukkan jumlah halaman : ")
	fmt.Scanln(&b.jumlahHalaman)

	fmt.Println("\n=== BIODATA BUKU ===")
	fmt.Println("Judul Buku :", b.judul)
	fmt.Println("Penulis :", b.penulis)
	fmt.Println("Penerbit :", b.penerbit)
	fmt.Println("Tahun Terbit :", b.tahunTerbit)
	fmt.Println("Jumlah Halaman :", b.jumlahHalaman)
}