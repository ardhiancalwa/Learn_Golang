package main

import "fmt"

func main() {
	var P, S, X, Y, B int
	var jumP, jumS int
	const hari int = 30

	fmt.Scan(&P, &S, &X, &Y, &B)
	for i := 1; i <= hari; i++ {
		if i%X != 0 && i%Y != 0 {
			if i%2 == 0 {
				jumS += S
			} else if i%2 != 0 {
				jumP += P
			}
		} 
	}
	jumP *= B
	jumS *= B

	fmt.Println(jumP, jumS)
}
