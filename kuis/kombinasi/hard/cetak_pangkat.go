package main

import (
	"fmt"
	"math"
)

func main() {
	var n, m int
	var i int

	fmt.Scan(&n, &m)
	for i = n; i <= m; i++ {
		if i%2 != 0 {
			fmt.Printf("%.2f", math.Pow(2, float64(i)))
			fmt.Print(" ")
		} else if i%2 == 0 {
			fmt.Printf("%.2f", math.Pow(2, float64(-i)))
			fmt.Print(" ")
		}
	}
}