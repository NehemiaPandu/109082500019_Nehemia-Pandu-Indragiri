package main

import "fmt"

import "math"

type titik struct {
	x, y float64
}

type Lingkaran struct {
	pusat  titik
	radius float64
}

func jarak(t, q titik) float64 {
	rumusJarak := math.Sqrt(math.Pow(t.x-q.x, 2) + math.Pow(t.y-q.y, 2))
	return rumusJarak
}

func diDalam(l Lingkaran, t titik) bool {

	return jarak(l.pusat, t) <= l.radius
}

func main() {
	var l1, l2 Lingkaran
	var t titik

	fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.radius)
	fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.radius)
	fmt.Scan(&t.x, &t.y)

	didalam1 := diDalam(l1, t)
	didalam2 := diDalam(l2, t)

	if didalam1 && didalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if didalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if didalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}