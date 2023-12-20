package main

import "fmt"

func main() {
	var kendaraan byte
	var bayar, jumBayarM, jumBayarB  int

	fmt.Scanf("%c", &kendaraan)

	if kendaraan != 'x' {
		if kendaraan == 'm' {
			fmt.Scan(&bayar)
			jumBayarM += bayar
		} else if kendaraan == 'b' {
			fmt.Scan(&bayar)
            jumBayarB += bayar
		}
		fmt.S
	}
	
}