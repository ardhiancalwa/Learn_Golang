package main
import "fmt"

func main()  {
	var n, totalJam, jam int
	var rata_rata float64
	
	n = 0
	totalJam = 0
	for {
		fmt.Scan(&jam)
		if jam < 0 {
			break
		}
		totalJam += jam
		n++
		rata_rata = float64(totalJam) / float64(n)
	}
	fmt.Printf("%.2f", rata_rata)
}
