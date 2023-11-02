package main
import "fmt"

func main()  {
	var bilangan, jumlah int

	jumlah = 0
	for {
		fmt.Scan(&bilangan)
		if bilangan < 0 {
			break
		}
		jumlah += bilangan
	}
	fmt.Println(jumlah)
	
}