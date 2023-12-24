package main

import "fmt"

func main() {
	var karakter byte
	var jumNumerik int
	
	for fmt.Scanf("%c", &karakter); karakter != '.'; fmt.Scanf("%c", &karakter) {
		if karakter >= '0' && karakter <= '9' {
			jumNumerik += int(karakter - '0')
		}
	}
	
	fmt.Println(jumNumerik)
}
