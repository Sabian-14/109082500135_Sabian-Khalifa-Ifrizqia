package main

import "fmt"

func ganjil(n int, i int) {
	if i > n {
		return
	}

	fmt.Printf("%d ", i)
	ganjil(n, i+2)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)

	fmt.Print("Bilangan ganjil: ")
	ganjil(n, 1)
}