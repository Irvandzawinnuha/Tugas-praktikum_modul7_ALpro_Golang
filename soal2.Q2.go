package main

import "fmt"

func average_1302220072(bil, i int, hasil *float64) {
	var dcount int
	if bil < 1 {
		*hasil = *hasil / float64(i)
		fmt.Println(*hasil)
	} else {
		dcount = bil % 10
		*hasil += float64(dcount)
		bil /= 10
		i++
		average_1302220072(bil, i, hasil)
	}
}

func main() {
	var res float64
	var xi, i int
	fmt.Println("masukkan angka tanpa spasi : ")
	fmt.Scan(&xi)
	average_1302220072(xi, i, &res)
}
