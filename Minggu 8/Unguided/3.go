package main

import "fmt"

func main() {
    var klubA, klubB string
    var skorA, skorB int
    var pemenang []string

    fmt.Print("Klub A : ")
    fmt.Scanln(&klubA)
    fmt.Print("Klub B : ")
    fmt.Scanln(&klubB)

    i := 1
    for {
        fmt.Printf("Pertandingan %d : ", i)
        fmt.Scan(&skorA, &skorB)

        if skorA < 0 || skorB < 0 {
            break
        }

        if skorA > skorB {
            pemenang = append(pemenang, klubA)
        } else if skorB > skorA {
            pemenang = append(pemenang, klubB)
        } else {
            pemenang = append(pemenang, "Draw")
        }
        i++
    }

    fmt.Println()
    for idx, hasil := range pemenang {
        fmt.Printf("Hasil %d : %s\n", idx+1, hasil)
    }
    fmt.Println("Pertandingan selesai")
}