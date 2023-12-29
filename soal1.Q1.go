package main

import "fmt"

func inputBilangan_1302223076(bil *int) {
	fmt.Scan(bil)
}

func stop(bil int) bool {
	return bil == 0
}

func hitung(mean *float64, min, max, n *int) {
	var bil, total int = 0, 0
	var p bool = true

	inputBilangan_1302223076(&bil)
	for !stop(bil) {
		if bil > 0 {
			total += bil
			*n++
			if bil > *max {
				*max = bil
			}
			if p || bil < *min {
				*min = bil
				p = false
			}
		}
		inputBilangan_1302223076(&bil)
	}
	*mean = float64(total) / float64(*n)
	if *n == 0 {
		fmt.Println("- - - -")
	} else {
		fmt.Println(*mean, *max, *min, *n)
	}
}

func main() {
	var mean float64
	var min, max, n int = 0, 0, 0

	hitung(&mean, &min, &max, &n)
}
