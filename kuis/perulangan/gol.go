package main
import "fmt"

func main()  {
	var jumGol, gol, i int
	jumGol = 0
	gol = 0
	fmt.Scan(&gol)
	for i = 1; i <= 30 && gol != -1; i++ {
		jumGol += gol 
		fmt.Scan(&gol)
	}
	fmt.Println(jumGol >= 10)
}	