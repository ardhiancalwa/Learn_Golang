package main

import "fmt"

func main() {
	var bil, digit, digitAfter int
	var status string

	fmt.Scan(&bil)
	status = "tidak terurut"

	for bil > 0 {
		digitAfter = digit
		digit = bil % 10

		if digitAfter != 0 {
			if digit < digitAfter {
				if status == "descending" {
					status = "tidak terurut"
					break
				}
				status = "ascending"
			} else if digit > digitAfter {
				if status == "ascending" {
					status = "tidak terurut"
					break
				}
				status = "descending"
			}
		}
		bil /= 10
	}

	fmt.Println(status)
}
