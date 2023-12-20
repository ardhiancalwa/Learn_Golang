package main

import "fmt"

func main() {
	var poin, i, n int
	var gold, platinum, silver int
	var newUser bool

	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		fmt.Scan(&poin, &newUser)
		for poin < 50 {
			fmt.Scan(&poin, &newUser)
		}
		if newUser {
			poin += 50
			if poin > 200 {
				gold++
			} else if poin >= 100 {
				platinum++
			} else if poin >= 50 {
				silver++
			}
		} else {
			if poin > 200 {
				gold++
			} else if poin >= 100 {
				platinum++
			} else if poin >= 50 {
				silver++
			}
		}
	}
	fmt.Println("gold: ", gold)
	fmt.Println("platinum: ", platinum)
	fmt.Println("silver: ", silver)
}
