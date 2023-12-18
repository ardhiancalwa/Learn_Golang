package main

import "fmt"

func main() {
	var i, n, lebar, max int

	fmt.Scan(&n)
	i = 1
	max = 0
	for i = 1; i <= n; i++ {
		fmt.Scan(&lebar)
		if i == 1 || lebar > max {
			max = lebar
		}
	}
	fmt.Println(max)
}