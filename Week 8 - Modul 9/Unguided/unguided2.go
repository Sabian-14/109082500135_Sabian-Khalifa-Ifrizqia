package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(&n)

	arr := make([]int, n)

	fmt.Println("Masukkan elemen array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("\nSemua elemen:")
	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println()

	fmt.Println("Indeks ganjil:")
	for i := 1; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Println("Indeks genap:")
	for i := 0; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	var x int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)

	fmt.Println("Indeks kelipatan", x, ":")
	for i := 0; i < len(arr); i++ {
		if x != 0 && i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	var del int
	fmt.Print("Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&del)

	if del >= 0 && del < len(arr) {
		arr = append(arr[:del], arr[del+1:]...)
	} else {
		fmt.Println("Index tidak valid, tidak dihapus")
	}

	fmt.Println("Setelah dihapus:")
	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println()

	sum := 0
	for _, v := range arr {
		sum += v
	}

	if len(arr) > 0 {
		avg := float64(sum) / float64(len(arr))
		fmt.Println("Rata-rata:", avg)

		var variance float64
		for _, v := range arr {
			diff := float64(v) - avg
			variance += diff * diff
		}
		variance /= float64(len(arr))
		stdDev := math.Sqrt(variance)

		fmt.Println("Standar deviasi:", stdDev)
	}

	var target int
	fmt.Print("Masukkan angka yang dicari: ")
	fmt.Scan(&target)

	count := 0
	for _, v := range arr {
		if v == target {
			count++
		}
	}

	fmt.Println("Frekuensi:", count)
}