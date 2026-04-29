package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var k rune
	fmt.Scanf("%c", &k)
	*n = 0
	for k != '.' && *n < NMAX {
		if k != ' ' && k != '\n' && k != '\r' {
			t[*n] = k
			(*n)++
		}
		fmt.Scanf("%c", &k)
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	var temp tabel
	temp = t
	balikanArray(&temp, n)
	for i := 0; i < n; i++ {
		if t[i] != temp[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	isiArray(&tab, &m)
	
	isPal := palindrom(tab, m)
	
	balikanArray(&tab, m)
	cetakArray(tab, m)
	
	fmt.Println(isPal)
}