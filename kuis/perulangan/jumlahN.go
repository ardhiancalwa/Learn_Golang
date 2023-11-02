package main
import "fmt"

func main()  {
	var i, n, bilangan, jumlah int

	fmt.Scan(&n)
	jumlah = 0
	i = 1
	for {
		fmt.Scan(&bilangan)
		if i > n  {
			break
		}
		jumlah += bilangan
		i++
	}
	fmt.Println(jumlah)
	
}