package main

import "fmt"

const pi float64 = 3.14

func volume (r, t float64) float64 {
	var volume float64

	volume = (pi * (r * r)) * t
	return  volume
}

func massa (r,t,p float64) float64 {
	var massa float64
	
	massa = volume(r,t) * p
	return massa
}

func display (m1, m2 float64) {
	if m1 == m2 {
		fmt.Println("\nBALANCE")
	} else {
		selisih := m1 - m2
		if selisih < 0 {
			selisih = -selisih
		}
		fmt.Print("\nSelisih massa zat cair kiri dan massa zat cair kanan: ", selisih)
	}
}

func main(){
	var r float64
	var tKiri, tKanan float64
	var mjKiri, mjKanan float64
	var massaKiri, massaKanan float64

	fmt.Print("Masukan jari - jari alas tabung: ")
	fmt.Scan(&r)

	fmt.Print("\nMasukan tinggi zat cair tabung kiri: ")
	fmt.Scan(&tKiri)
	fmt.Print("Masukan massa jenis zat cair tabung kiri: ")
	fmt.Scan(&mjKiri)
	
	fmt.Print("\nMasukan tinggi zat cair tabung kanan: ")
	fmt.Scan(&tKanan)
	fmt.Print("Masukan massa jenis zat cair tabung kanan: ")
	fmt.Scan(&mjKanan)
	
	massaKiri = massa(r, tKiri, mjKiri)
	massaKanan = massa(r, tKanan, mjKanan)

	display(massaKiri, massaKanan)
}