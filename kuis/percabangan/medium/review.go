package main
import "fmt"

func main()  {
	var review1, review2, review3, review4 int
	var bagus float64 = 3.50
	var kurang float64 = 1.50
	fmt.Scan(&review1, &review2, &review3, &review4)
	if  (review1 + review2 + review3 + review4) / 4  >= int(bagus){
		fmt.Println("bagus")
	} else if (review1 + review2 + review3 + review4) / 4 <= int(kurang){
		fmt.Println("kurang")
	} else {
		fmt.Println("sedang")
	}

}