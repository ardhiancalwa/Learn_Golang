package main
import "fmt"

func main()  {
	var i, jam, menit, jumJam, jumMenit int

	i = 1
	for {
		fmt.Scan(&jam, &menit)
		jumJam += jam
		jumMenit += menit 
		jumJam += jumMenit / 60
		jumMenit %= 60
		i++
		if jam == 0 && menit == 0 {
			break
		}
	}
	fmt.Println(jumJam, jumMenit)
}