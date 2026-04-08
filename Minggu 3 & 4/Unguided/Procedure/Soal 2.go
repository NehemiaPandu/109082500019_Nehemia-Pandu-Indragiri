package main

import "fmt"

func hitungPersegi (sisi int) {
	var luas, keliling int

	luas = sisi * sisi
	keliling = sisi * 4

	fmt.Println("Luas Persegi:", luas)
	fmt.Println("Keliling Persegi:", keliling)
}

func hitungPersegiPanjang (panjang, lebar int){
	var luas, keliling int
	
	luas = panjang * lebar
	keliling = (panjang + lebar) * 2
	
	fmt.Println("Luas Persegi Panjang:", luas)
	fmt.Println("Keliling Persegi Panjang:", keliling)
}

func hitungLingkaran (jariJari float64) {
	var luas, keliling, pi float64
	
	pi = 3.14
	
	luas = pi * jariJari * jariJari
	keliling = 2 * pi * jariJari
	
	fmt.Println("Luas Lingkaran:", luas)
	fmt.Println("Keliling Lingkaran:", keliling)
}

func main (){
	var pilihan int

	fmt.Println("--- PROGRAM BANGUN DATAR ---")
	fmt.Println("1. Hitung luas & keliling persegi")
	fmt.Println("2. Hitung luas & keliling persegi panjang")
	fmt.Println("3. Hitung luas & keliling lingkaran")
	fmt.Print("Pilihan: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		var s int
		fmt.Print("Masukan sisi: ")
		fmt.Scan(&s)
		hitungPersegi(s)
	case 2:
		var p, l int
		fmt.Print("Masukan panjang: ")
		fmt.Scan(&p)
		fmt.Print("Masukan lebar: ")
		fmt.Scan(&l)
		hitungPersegiPanjang(p, l)
	case 3: 
		var r float64
		fmt.Print("Masukan jari-jari: ")
		fmt.Scan(&r)
		hitungLingkaran(r)
	}

}