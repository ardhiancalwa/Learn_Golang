package main

import "fmt"

func main() {
	var P, S, X, Y, B int
	var jumP, jumS int
	const hari int = 30

	fmt.Scan(&P, &S, &X, &Y, &B)
	for j := 1; j <= hari; j++ {
		if j%X == 0 || j%Y == 0 {
			jumP += 0
			jumS += 0
		} else {
			if j%2 == 0 {
				jumS += S
			} else {
				jumP += P
			}
		}
	}
	jumP *= B
	jumS *= B

	fmt.Println(jumP, jumS)
}
