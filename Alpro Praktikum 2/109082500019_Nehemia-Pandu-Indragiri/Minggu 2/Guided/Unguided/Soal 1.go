package main

import "fmt"

func main(){
	var tahun int
	var kabisat bool

	fmt.Print("Tahun: ")
	fmt.Scan(&tahun)

	if tahun % 400 == 0 || tahun % 4 == 0 && tahun % 100 != 0 {
		kabisat = true
		fmt.Println("Kabisat:", kabisat)
		} else {
		kabisat = false
		fmt.Println("Kabisat:", kabisat)
		
	}

}