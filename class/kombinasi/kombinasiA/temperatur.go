package main

import "fmt"

func main() {
	var suhu int
	var tertinggi, terendah, n int
	var rata float64

	fmt.Scan(&suhu)
	for suhu != 0 {
		n++
		if suhu > tertinggi {
			tertinggi = suhu
		}
		if suhu < terendah {
			terendah = suhu
		}
		rata += float64(suhu)
		rata /= float64(n)
	}
	fmt.Println(tertinggi, terendah, rata)
}