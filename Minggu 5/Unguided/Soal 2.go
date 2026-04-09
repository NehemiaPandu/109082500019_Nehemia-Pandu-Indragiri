package main

import "fmt"

func cetakAngka(n int) {
	if n <= 0 {
		return
	}

	fmt.Print(n, " ")

	if n > 1 {
		cetakAngka(n - 1)

		fmt.Print(n, " ")
	}
}

func main() {
	var input int

	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&input)

	cetakAngka(input)
	fmt.Println()
}