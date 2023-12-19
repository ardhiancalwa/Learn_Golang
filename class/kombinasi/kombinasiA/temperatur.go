package main

import "fmt"

func main() {
	var suhu, terendah, tertinggi, stop, n, rata int
	var jumRata float64

	for {
		fmt.Scan(&suhu)
		if suhu == 0 && stop == 0 {
			break
		}
		n++
		if n == 1 {
			tertinggi = suhu
			terendah = suhu
		} else if suhu < terendah {
			terendah = suhu
		} else if suhu > tertinggi {
			tertinggi = suhu
		}

		rata += suhu
		jumRata = float64(rata) / float64(n)
		stop = suhu
	}
	fmt.Println(tertinggi, terendah, jumRata)
}
