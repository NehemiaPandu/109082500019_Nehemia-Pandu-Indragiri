package main

import "fmt"

type Wadah struct {
	totalBanyakIkan float64
}

func main() {
	var x, y int
	var dataIkan [1000]float64
	var daftarWadah [1000]Wadah

	fmt.Print("Masukkan x (jumlah data) dan y (kapasitas wadah): ")
	fmt.Scan(&x, &y)

	fmt.Printf("Masukkan %d data nilai: ", x)
	for i := 0; i < x; i++ {
		fmt.Scan(&dataIkan[i])
	}

	var tempSum float64
	var tempCount int
	var nWadah int = 0

	for i := 0; i < x; i++ {
		tempSum += dataIkan[i]
		tempCount++

		if tempCount == y || i == x-1 {
			daftarWadah[nWadah].totalBanyakIkan = tempSum
			tempSum = 0
			tempCount = 0
			nWadah++
		}
	}

	fmt.Println("\n--- HASIL ---")
	
	fmt.Print("Total per wadah: ")
	for i := 0; i < nWadah; i++ {
		fmt.Printf("%.2f ", daftarWadah[i].totalBanyakIkan)
	}
	fmt.Println()

	var totalSemua float64
	for i := 0; i < nWadah; i++ {
		totalSemua += daftarWadah[i].totalBanyakIkan
	}
	
	if nWadah > 0 {
		fmt.Printf("Rata-rata wadah: %.2f\n", totalSemua/float64(nWadah))
	}
}