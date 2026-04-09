package main

import "fmt"

func banyakBintang(jumlah int) {
	if jumlah <= 0 {
		return
	}
	fmt.Print("*")
	banyakBintang(jumlah - 1)
}

func baris(n int) {
	if n <= 0 {
		return
	}
	baris(n - 1)

	banyakBintang(n)
	fmt.Println()
}

func main() {
	var input int

	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&input)

	baris(input)
}
