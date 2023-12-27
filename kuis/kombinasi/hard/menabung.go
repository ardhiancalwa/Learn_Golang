package main

import "fmt"

func main() {
	var x, y, z, y2 int
	var totTabungan int
	var i int

	fmt.Scan(&x, &y, &z)
	for i = 1; i <= z; i++ {
		if i%2 == 0 || i%3 == 0 {
			if totTabungan == 0 {
				totTabungan = x
			} else {
				if y2 > 2 {
					y += y
					totTabungan += x + y
				} else {
					totTabungan += x + y
				}
			}
		}
		y2++
	}
	fmt.Println(totTabungan)
}