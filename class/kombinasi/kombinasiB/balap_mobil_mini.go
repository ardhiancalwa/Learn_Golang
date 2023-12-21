package main

import "fmt"

func main() {
	var A, B, i int
	var tim byte

	for i = 1; i <= 10; i++ {
		fmt.Scanf("%c", &tim)
		if tim == 'A'{
			A++
		}
		if tim == 'B'{
            B++
        }
		if A == 4 || B == 4 {
			break
		}
		fmt.Scanf("%c", &tim)
	}
	fmt.Println(A, B)
	if A == 4 {
		fmt.Println("A")	
	} else if B == 4 {
		fmt.Println("B")
	} else if A >= 5 || B >= 5 {
		fmt.Println("tidak valid")
	} 
}