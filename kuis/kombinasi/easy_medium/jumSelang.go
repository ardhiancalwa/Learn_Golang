package main

import "fmt"

func main() {
	var i, n, m int
	var jumDeret float64

	fmt.Scan(&n, &m)
	if n < m && n > 0{
		for i = n; i <= m; i++ {
			if i%2 != 0 {
				jumDeret += 3.0 /  float64(i)
			} else {
				jumDeret -= 3.0 / float64(i)
			}
		}
		fmt.Printf("%.2f \n", jumDeret)
	}
}
