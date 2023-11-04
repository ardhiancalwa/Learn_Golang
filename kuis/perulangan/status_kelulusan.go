package main
import "fmt"

func main() {
	var clo1, clo2, clo3 float64
	var lulus bool

	for {
		fmt.Scan(&clo1, &clo2, &clo3)
		lulus = true
		lulus = lulus && clo1 > 50 && clo2 > 50 && clo3 > 50

		if clo1 < 0 || clo2 < 0 || clo3 < 0 {
			break
		}

		fmt.Println(lulus)
	}
}
