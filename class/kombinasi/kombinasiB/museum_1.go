package main

import "fmt"

func main() {
	var pengunjung, jumPengunjung, kenaikan, status, hari int
	var rata float64

	fmt.Scan(&pengunjung)
	for pengunjung > 0 && pengunjung <= 200 {
		hari++
		jumPengunjung += pengunjung
		if hari > 1 && status < pengunjung {
			kenaikan++
		}
		status = pengunjung
		fmt.Scan(&pengunjung)
	}
	if hari == 0 {
		hari = 1
	}
	rata = float64(jumPengunjung)/float64(hari)
	fmt.Println(kenaikan, rata)
}