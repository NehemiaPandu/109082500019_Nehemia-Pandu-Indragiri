package main

import "fmt"

const nProv = 10

type NamaProv [nProv]string
type PopProv [nProv]int
type TumbuhProv [nProv]float64

func main() {
	var provinsi NamaProv
	var populasi PopProv
	var pertumbuhan TumbuhProv
	var cariNama string

	fmt.Println("=== Masukkan Nama Provinsi, Populasi Provinsi, Angka Pertumbuhan Provinsi ===")
	InputData(&provinsi, &populasi, &pertumbuhan)
	fmt.Scan(&cariNama)

	fmt.Println()
	idxCepat := ProvinsiTercepat(pertumbuhan)
	fmt.Printf("Provinsi dengan angka pertumbuhan tercepat : %s\n", provinsi[idxCepat-1])

	fmt.Println()
	fmt.Printf("Data provinsi yang dicari : %s\n", cariNama)
	
	fmt.Println()
	fmt.Println("=== Prediksi Jumlah Penduduk Tahun Depan Pada Provinsi Dengan Pertumbuhan Diatas 2% ===")
	Prediksi(provinsi, populasi, pertumbuhan)
}

func InputData(prov *NamaProv, pop *PopProv, tumbuh *TumbuhProv) {
	for i := 0; i < nProv; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&prov[i], &pop[i], &tumbuh[i])
	}
}

func ProvinsiTercepat(tumbuh TumbuhProv) int {
	maxIdx := 0
	for i := 1; i < nProv; i++ {
		if tumbuh[i] > tumbuh[maxIdx] {
			maxIdx = i
		}
	}
	return maxIdx + 1
}

func IndeksProvinsi(prov NamaProv, nama string) int {
	for i := 0; i < nProv; i++ {
		if prov[i] == nama {
			return i + 1
		}
	}
	return -1
}

func Prediksi(prov NamaProv, pop PopProv, tumbuh TumbuhProv) {
	for i := 0; i < nProv; i++ {
		if tumbuh[i] > 0.02 {
			prediksi := (tumbuh[i] + 1.0) * float64(pop[i])
			fmt.Printf("%s %.0f\n", prov[i], prediksi)
		}
	}
}