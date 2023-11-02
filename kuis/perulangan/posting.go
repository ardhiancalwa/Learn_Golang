package main
import "fmt"

func main()  {
	var i, n, hari, postingan, jumHari, jumPost int

	fmt.Scan(&n)
	jumHari = 0
	jumPost = 0
	i = 1
	for i <= n {
		fmt.Scan(&hari)
		jumHari += hari
		fmt.Scan(&postingan)
		jumPost += postingan
		i++
	}
	fmt.Println(jumHari, jumPost)
}