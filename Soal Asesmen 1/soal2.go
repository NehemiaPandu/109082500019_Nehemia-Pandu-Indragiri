package main

import "fmt"

func tanggunganHari(jumlahHari int, tujuan string) int {
	if tujuan == "domestik"{
		if jumlahHari > 3 {
			return 3
		}
	} else {
		if jumlahHari > 8 {
			return 8
		}
	}
	return jumlahHari
}

func biayaPerHari(jumlahMahasiswa int) int {
	biaya := (2 * 35000) + 250000 + 300000
		return biaya * jumlahMahasiswa	
}

func perhitunganBiaya(jumlahMahasiswa, lamaPerjalanan int, tujuan string, totalBiaya *float64) {
	hari := tanggunganHari(lamaPerjalanan, tujuan)

	totalPerHari := biayaPerHari(jumlahMahasiswa)

	if tujuan == "mancanegara" {
		totalPerHari = int(float64(totalPerHari) * 1.5)
	}

	*totalBiaya = float64(hari * totalPerHari)
}

func main(){
	var jumlah, lama int
	var tujuan string
	var biaya float64

	fmt.Print("Masukan jumlah mahasiswa: ")
	fmt.Scan(&jumlah)

	fmt.Print("Masukan lama hari study tour: ")
	fmt.Scan(&lama)

	fmt.Print("Masukan tujuan study tour (domestik/mancanegara): ")
	fmt.Scan(&tujuan)

	perhitunganBiaya(jumlah, lama, tujuan, &biaya)

	fmt.Printf("\nBiaya yang harus dikeluarkan Tel-U: %.0f\n", biaya)
}