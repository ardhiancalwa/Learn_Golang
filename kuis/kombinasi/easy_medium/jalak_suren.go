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
        } 
		if suara == "BE" {
			BE = true
		} 
		if suara == "CI" {
			CI = true
		} 
		if suara == "JE" {
			JE = true
		} 
		if suara == "LO" {
			LO = true
		} 
	}

	fmt.Println(JA, BE, CI, JE, LO)

	if JA && BE && CI && JE && LO {
		lima_suara = true
	} else {
		lima_suara = false
	}
	fmt.Println(lima_suara)
}