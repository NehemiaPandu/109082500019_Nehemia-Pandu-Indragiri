	package main

	import "fmt"

	const NMAX = 1000

	type arrInt [NMAX]int

	func SelectionSort(T *arrInt, n int) {
		for i := 0; i < n-1; i++ {
			idxMin := i
			for j := i + 1; j < n; j++ {
				if T[j] < T[idxMin] {
					idxMin = j
				}
			}
			T[i], T[idxMin] = T[idxMin], T[i]
		}
	}

	func median(T arrInt, n int) float64 {
		if n%2 != 0 {
			return float64(T[n/2])
		}
		return float64(T[n/2-1]+T[n/2]) / 2.0
	}

	func main() {
		var A arrInt
		var x int
		var n int = 0

		fmt.Scan(&x)
		for x != -5313541 && n < NMAX {
			if x == 0 {
				SelectionSort(&A, n)
				hasilMedian := median(A, n)
				if hasilMedian == float64(int(hasilMedian)) {
					fmt.Println(int(hasilMedian))
				} else {
					fmt.Println(hasilMedian)
				}
			} else {
				A[n] = x
				n++
			}
			fmt.Scan(&x)
		}
	}