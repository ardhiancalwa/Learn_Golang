package main

import "fmt"

func main() {
	var poin int
	var newUser bool

	fmt.Scan(&poin, &newUser)
	for poin < 50 {
		fmt.Scan(&poin, &newUser)
	}
	if newUser {
		poin += 50
		if poin > 200 {
			fmt.Println("Gold User")
		} else if poin >= 100 {
			fmt.Println("Platinum User")
		} else if poin >= 50 {
			fmt.Println("Silver User")
		}
	} else {
		if poin > 200 {
			fmt.Println("Gold User")
		} else if poin >= 100 {
			fmt.Println("Platinum User")
		} else if poin >= 50 {
			fmt.Println("Silver User")
		}
	}
}
