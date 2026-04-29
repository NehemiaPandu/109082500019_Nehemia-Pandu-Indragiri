package main

import (
	"fmt"
	"math"
)

func main() {
	var n, x, indexHapus int
	fmt.Print("Masukkan jumlah elemen (N): ")
	fmt.Scan(&n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Isi indeks ke-%d: ", i)
		fmt.Scan(&arr[i])
	}

	fmt.Println("\n--- LAPORAN DATA ---")

	fmt.Print("a. Seluruh isi: ")
	fmt.Println(arr)

	fmt.Print("b. Indeks ganjil: ")
	for i := 1; i < len(arr); i += 2 {
		fmt.Printf("%d ", arr[i])
	}

	fmt.Print("\nc. Indeks genap: ")
	for i := 0; i < len(arr); i += 2 {
		fmt.Printf("%d ", arr[i])
	}

	fmt.Print("\nd. Masukkan nilai x: ")
	fmt.Scan(&x)
	fmt.Print("   Elemen kelipatan x: ")
	for i := 0; i < len(arr); i++ {
		if i%x == 0 {
			fmt.Printf("%d ", arr[i])
		}
	}

	total := 0
	for _, v := range arr {
		total += v
	}
	rataRata := float64(total) / float64(len(arr))
	fmt.Printf("\nf. Rata-rata: %.2f", rataRata)

	var varians float64
	for _, v := range arr {
		varians += math.Pow(float64(v)-rataRata, 2)
	}
	sd := math.Sqrt(varians / float64(len(arr)))
	fmt.Printf("\ng. Standar Deviasi: %.2f", sd)

	fmt.Print("\ne. Hapus indeks ke: ")
	fmt.Scan(&indexHapus)
	
	arr = append(arr[:indexHapus], arr[indexHapus+1:]...)
	
	fmt.Print("   Isi array setelah dihapus: ")
	fmt.Println(arr)
}