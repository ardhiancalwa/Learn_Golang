package main
import "fmt"

func main()  {
	var kapasitas, jumPenumpang float64

	fmt.Scan(&kapasitas, &jumPenumpang)
	if jumPenumpang >= 0.5 * kapasitas && jumPenumpang <= 0.75 * kapasitas {
		fmt.Println("berangkat")
	} else {
		fmt.Println("tidak berangkat")
	}
}