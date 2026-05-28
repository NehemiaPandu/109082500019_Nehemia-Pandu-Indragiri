package main

import "fmt"

func SequentialSearch(arrBuah [5] string, dataDicari string) int {
	var idx_found = -1 
	for i:=0; i < len(arrBuah); i++ {
		if arrBuah[i] == dataDicari {
			idx_found = i
			break
		}
	}
	return idx_found
}

func main(){
	var arrBuah [5] string

	for i := 0; i < len(arrBuah); i++ {
		fmt.Printf("Masukan data buah indeks ke - %d : ", i)
		fmt.Scan(&arrBuah[i])
	}

	var dataCari string
	fmt.Println("Masukan data buah yang ingin dicari: ")
	fmt.Scan(&dataCari)

	var indexData int
	indexData = SequentialSearch(arrBuah, dataCari)

	if indexData > -1 {
		fmt.Printf("Data %s ditemukan pada indeks ke-%d!", dataCari, indexData)
	}else if indexData == -1 {
		fmt.Printf("Data %s tidak ditemukan", dataCari)
	}

}
