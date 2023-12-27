package main

import "fmt"

func main() {
	var prima bool = true
	var bil, i int

	fmt.Scan(&bil)
	if bil <= 1 {
		prima = false
	} else {
		for i = 2; i*i <= bil; i++ {
			if bil%i == 0 {
                prima = false
            }
        }
	}
	fmt.Println(prima)
}