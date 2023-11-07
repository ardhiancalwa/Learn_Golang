package main
import "fmt"

func main()  {
	var beasiswa string
	
	fmt.Scan(&beasiswa)
	for beasiswa != "ada" {
		fmt.Println("cari beasiswa")
		fmt.Scan(&beasiswa)
	}
	fmt.Println("pencarian selesai")
}