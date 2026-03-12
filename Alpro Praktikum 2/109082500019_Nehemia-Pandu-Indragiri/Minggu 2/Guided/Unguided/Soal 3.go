package main

import "fmt"

func main(){
	var Totalberat, kg, gram, jasaKg, jasaGram int

	fmt.Print("Masukan total berat (gram): ")
	fmt.Scan(&Totalberat)
	fmt.Println()

	fmt.Println("===== Detail Perhitungan =====")
	kg = Totalberat / 1000
	gram = Totalberat % 1000

	fmt.Println("Detail Berat:", kg, "kg +", gram, "gram")
	
	jasaKg = kg * 10000
	
	if kg > 10 {
		jasaGram = 0
	}else{
		if gram >= 500 {
			jasaGram = 5 * gram
		} else if gram < 500 {
		jasaGram = 15 * gram
		}
	}

	fmt.Println("Detail Biaya: Rp.", jasaKg, "+ Rp.", jasaGram)
	fmt.Println("Total Biaya: Rp.", jasaKg + jasaGram)

}