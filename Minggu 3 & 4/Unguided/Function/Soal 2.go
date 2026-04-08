package main

import "fmt"

func tarifMotor (jam int) int {
	if jam >= 0 && jam < 17 {
		return 4000
	} else {
		return 5000
	}	
}

func tarifMobil (jam int) int {
	if jam >= 0 && jam < 17 {
		return 6000
	} else {
		return 7000
	}
}

func main (){
	fmt.Println("=== Rekap Tarif Parkir Cafe per Hari ===")

	var i int = 1
	var jenisKendaraan string
	var jamMasuk int
	var jamKeluar int
	var pendapatan int = 0
	

	for {
		fmt.Println("\n*Kendaraan", i)

		fmt.Print("Kendaraan apa? (motor/mobil/- untuk selesai): ")
		fmt.Scan(&jenisKendaraan)
		
		if jenisKendaraan == "-" {
			break
		}

		fmt.Print("Masukan Jam Masuk Kendaraan (0-24): ")
		fmt.Scan(&jamMasuk)
		fmt.Print("Masukan Jam Keluar Kendaraan (0-24): ")
		fmt.Scan(&jamKeluar)

		total := 0

		for j := jamMasuk; j < jamKeluar; j++ {
			if jenisKendaraan == "motor" {
				total += tarifMotor(j)
			} else if jenisKendaraan == "mobil" {
				total += tarifMobil(j)
			}
		}

		if jenisKendaraan == "motor" {
			fmt.Println("Biaya parkir motor", i, ": ", total)
		} else if jenisKendaraan == "mobil" {
			fmt.Println("Biaya parkir mobbil", i, ": ", total)
		}

		fmt.Println("========================================")

		pendapatan += total
		i++
		
		
	}
	
	fmt.Println("\n*** Total Pendapatan Parkir Hari Ini Adalah", pendapatan, "***")
}