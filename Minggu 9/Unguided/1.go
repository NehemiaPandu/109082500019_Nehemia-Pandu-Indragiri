package main

import "fmt"

func main() {
    var berat [1000]float64
    var n int

    fmt.Print("Masukan banyak data berat anak kelinci: ")
    fmt.Scan(&n)

    if n > 1000 {
        n = 1000
    }

    for i := 0; i < n; i++ {
        fmt.Printf("Masukan berat kelinci ke-%d: ", i+1)
        fmt.Scan(&berat[i])
    }

    var min, max float64
    if n > 0 {
        min = berat[0]
        max = berat[0]
    }

    for i := 1; i < n; i++ {
        if berat[i] < min {
            min = berat[i]
        }
        if berat[i] > max {
            max = berat[i]
        }
    }

    fmt.Printf("\nBerat kelinci terkecil: %.2f\n", min)
    fmt.Printf("Berat kelinci terbesar: %.2f\n", max)
}