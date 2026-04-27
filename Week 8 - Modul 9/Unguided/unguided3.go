package main

import "fmt"

func main() {
	var klubA, klubB string
	fmt.Print("klub A: ")
	fmt.Scan(&klubA)

	fmt.Print("klub B: ")
	fmt.Scan(&klubB)

	winners := []string{}
	match := 1

	for {
		var scoreA, scoreB int
		fmt.Printf("Pertandingan %d : ", match)
		fmt.Scan(&scoreA, &scoreB)

		if scoreA < 0 || scoreB < 0 {
			break
		}

		if scoreA > scoreB {
			fmt.Printf("Hasil %d : %s\n", match, klubA)
			winners = append(winners, klubA)
		} else if scoreB > scoreA {
			fmt.Printf("Hasil %d : %s\n", match, klubB)
			winners = append(winners, klubB)
		} else {
			fmt.Printf("Hasil %d : Draw\n", match)
		}

		match++
	}

	fmt.Println("Pertandingan selesai")

	fmt.Println("Daftar pemenang:")
	for i, w := range winners {
		fmt.Printf("%d. %s\n", i+1, w)
	}
}