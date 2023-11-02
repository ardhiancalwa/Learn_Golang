package main
import "fmt"

func main()  {
	var beasiswa string
	
	for {
		fmt.Scan(&beasiswa)
		if beasiswa == "belum" {
			fmt.Println("cari beasiswa")
		} else if beasiswa == "ada" {
			fmt.Println("pencarian selesai")
			break
		}
	}
}