package main

import "fmt"

func main() {
	var bil1, bil2 int
	var jumBil1, jumBil2, totBil int

	fmt.Scan(&bil1, &bil2)
	for bil1 != 0 && bil2 != 0 {
		if bil1%2 != 0 && bil2%2 == 0 {
			jumBil1 += bil1
			jumBil2 += bil2
		}
		fmt.Scan(&bil1, &bil2)
	}
	totBil = jumBil1 + jumBil2
	fmt.Println(totBil)
}
