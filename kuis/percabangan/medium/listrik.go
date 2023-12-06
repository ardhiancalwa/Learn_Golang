package main
import "fmt"

func main() {
	var unit, biaya, jumBayar int
	var tambahan float64

	fmt.Scan(&unit)
	if unit < 200 {
		biaya = 12
			if unit*biaya > 400 {
				tambahan = float64(0.2) * float64(unit) * float64(biaya)
				jumBayar = (unit * biaya) + int(tambahan)
			} else if unit * biaya < 100 {
				jumBayar = 100
			} else {
				jumBayar = unit * biaya
			}
	} else if unit >= 200 && unit < 400 {
		biaya = 15
			if unit*biaya > 400 {
				tambahan = float64(0.2) * float64(unit) * float64(biaya)
				jumBayar = (unit * biaya) + int(tambahan)
			} else if unit * biaya < 100 {
				jumBayar = 100
			} else {
				jumBayar = unit * biaya
			}
	} else if unit >= 400 && unit < 600 {
		biaya = 18
			if unit*biaya > 400 {
				tambahan = float64(0.2) * float64(unit) * float64(biaya)
				jumBayar = (unit * biaya) + int(tambahan)
			} else if unit * biaya < 100 {
				jumBayar = 100
			} else {
				jumBayar = unit * biaya
			}
	} else {
		biaya = 20
			if unit*biaya > 400 {
				tambahan = float64(0.2) * float64(unit) * float64(biaya)
				jumBayar = (unit * biaya) + int(tambahan)
			} else if unit * biaya < 100 {
				jumBayar = 100
			} else {
				jumBayar = unit * biaya
			}
	}
	fmt.Println(jumBayar)
}
