package main

import "fmt"

type Pemain struct {
	NamaDepan    string
	NamaBelakang string
	Gol          int
	Assist       int
}

func lebihBaik(p1, p2 Pemain) bool {
	if p1.Gol > p2.Gol {
		return true
	} else if p1.Gol == p2.Gol {
		if p1.Assist > p2.Assist {
			return true
		}
	}
	return false
}

func insertionSort(arr []Pemain, n int) {
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && lebihBaik(key, arr[j]) {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	var daftarPemain [1005]Pemain

	for i := 0; i < n; i++ {
		fmt.Scan(&daftarPemain[i].NamaDepan, &daftarPemain[i].NamaBelakang, &daftarPemain[i].Gol, &daftarPemain[i].Assist)
	}

	insertionSort(daftarPemain[:], n)

	fmt.Println("---------------------")

	for i := 0; i < n; i++ {
		fmt.Printf("%s %s %d %d\n", daftarPemain[i].NamaDepan, daftarPemain[i].NamaBelakang, daftarPemain[i].Gol, daftarPemain[i].Assist)
	}
}