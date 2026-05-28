package main

import "fmt"

type Calon struct {
	ID    int
	Suara int
}

func main() {
	var num int
	var suaraMasuk, suaraSah int
	var daftarCalon [20]Calon

	for i := 0; i < 20; i++ {
		daftarCalon[i].ID = i + 1
		daftarCalon[i].Suara = 0
	}

	for {
		_, err := fmt.Scan(&num)
		if err != nil || num == 0 {
			break
		}

		suaraMasuk++

		if num >= 1 && num <= 20 {
			suaraSah++

			for i := 0; i < 20; i++ {
				if daftarCalon[i].ID == num {
					daftarCalon[i].Suara++
					break
				}
			}
		}
	}

	fmt.Printf("Suara masuk: %d\n", suaraMasuk)
	fmt.Printf("Suara sah: %d\n", suaraSah)

	for i := 0; i < 20; i++ {
		if daftarCalon[i].Suara > 0 {
			fmt.Printf("%d: %d\n", daftarCalon[i].ID, daftarCalon[i].Suara)
		}
	}
}