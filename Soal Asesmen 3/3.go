package main

import "fmt"

const NMAX = 1000000

type partai struct {
	nama  int
	suara int
}

type tabPartai [NMAX]partai

func main() {
	var p tabPartai
	var nPartai int = 0
	var suaraInput int

	fmt.Println("Masukkan proses input suara :")
	for {
		fmt.Scan(&suaraInput)
		if suaraInput == -1 {
			break
		}

		idx := posisi(p, nPartai, suaraInput)
		
		if idx != -1 {
			p[idx].suara++
		} else {
			p[nPartai].nama = suaraInput
			p[nPartai].suara = 1
			nPartai++
		}
	}

	for i := 1; i < nPartai; i++ {
		key := p[i]
		j := i - 1

		for j >= 0 && p[j].suara < key.suara {
			p[j+1] = p[j]
			j--
		}
		p[j+1] = key
	}

	fmt.Println()
	fmt.Println("Hasil Perhitungan suara :")
	for i := 0; i < nPartai; i++ {
		fmt.Printf("%d(%d) ", p[i].nama, p[i].suara)
	}
	fmt.Println()
}

func posisi(t tabPartai, n int, nama int) int {
	for i := 0; i < n; i++ {
		if t[i].nama == nama {
			return i
		}
	}
	return -1
}