package main

import "fmt"

func hitungPangkat(x, y int) int {
	if y == 0 {
		return 1
	}
	
	return x * hitungPangkat(x, y-1)
}

func main() {
	var x, y int

	fmt.Print("Masukkan nilai x dan y: ")
	fmt.Scan(&x, &y)

	hasil := hitungPangkat(x, y)
	
	fmt.Println("Hasil:", hasil)
}