package main

import "fmt"

type suhu float64

func CelciusToReamur(Celcius suhu) suhu{
	reamur := 4.0/5.0 * Celcius
	return reamur
}

func CelciusToFahrenheit (Celcius suhu) suhu{
	fahrenheit := (9.0/5.0 * Celcius) + 32
	return fahrenheit
}

func CelciusToKelvin (Celcius suhu) suhu{
	kelvin := Celcius + 273.15
	return kelvin
}

func main () {
	var c suhu

	fmt.Println("===== KONVERTER CELCIUS =====")
	fmt.Print("Masukan Suhu (Celcius): ")
	fmt.Scan(&c)
	fmt.Println()

	fmt.Println(c, "celcius =", CelciusToReamur(c), "reamur")
	fmt.Println(c, "celcius =", CelciusToFahrenheit(c), "fahrenheit")
	fmt.Println(c, "celcius =", CelciusToKelvin(c), "kelvin")

}