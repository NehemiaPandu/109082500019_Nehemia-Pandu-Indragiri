package main

import "fmt"

func main() {
	var a, b, c, d string
	var berhasil bool

	for i := 1; i <= 5; i++ {
		fmt.Print("Percobaan ke-", i, ": ")
		fmt.Scan(&a, &b, &c, &d)

		if a == "merah" && b == "kuning" && c == "hijau" && d == "ungu" {
			berhasil = true
		} else {
			berhasil = false
		}
	}

	fmt.Println("Berhasil: ", berhasil)

}
