package main
import "fmt"

func main() {
	var i, bilangan, bilanganAkhir, banyak int

	i = 0
	for i < 10 {
		fmt.Scan(&bilangan)
		if i == 0 || bilangan <= bilanganAkhir {
			bilanganAkhir = bilangan
			banyak++
		} else {
			break
		}
		i++
	}

	fmt.Println(banyak)
}
