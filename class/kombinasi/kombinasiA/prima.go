package main

import "fmt"

func main() {
	var bil, i int
	var prima bool = true

	fmt.Scan(&bil)
	if bil <= 1 {
		prima = false
	} else {
		i = 2
		for i*i <= bil {
			if bil%i == 0 {
				prima = false
			}
			i++
		}
	}
	fmt.Println(prima)
}
