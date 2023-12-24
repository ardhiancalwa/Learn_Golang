package main

import "fmt"

func main() {
	var suara string
	var i, n int
	var lima_suara bool
	var JA, BE, CI, JE, LO bool

	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		fmt.Scan(&suara)
		if suara == "JA" {
            JA = true
        } else {
			JA = false
		}
		if suara == "BE" {
			BE = true
		} else {
			BE = false
		}
		if suara == "CI" {
			CI = true
		} else {
			CI = false
		}
		if suara == "JE" {
			JE = true
		} else {
			JE = false
		}
		if suara == "LO" {
			LO = true
		} else {
			LO = false
		}
	}

	if JA == true && BE == true && CI == true && JE == true && LO == true {
		lima_suara = true
	} else {
		lima_suara = false
	}
	fmt.Println(lima_suara)
}