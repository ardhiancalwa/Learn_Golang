package main
import "fmt"
func main()  {
	var a, b, c, d, u1, u2, u3, u4 int
	var maxAB, minAB, maxCD, minCD int
	var temp int
	var mean, median float64

	fmt.Scan(&a, &b, &c, &d)
	mean = float64(a+b+c+d)/4
	if a >= b {
		maxAB = a
		minAB = b
	} else {
		maxAB = b
		minAB = a
	}
	if c >= d {
		maxCD = c
		maxCD = d
	} else {
		maxCD = d
		maxCD = c
	}
	if maxAB >= maxCD {
		u4 = maxAB
		u3 = maxCD
	} else {
		u4 = maxCD
		u3 = maxAB
	}
	if minAB <= minCD {
		u1 = minAB
		u2 = minCD
	} else {
		u1 = minCD
		u2 = minAB
	}
	if u3 <= u2 {
		temp = u3
		u3 = u2
		u2 = temp
	}
	median = float64(u2+u3)/2
	fmt.Println(u1, u4)
	fmt.Println(u1, u2, u3, u4)
	fmt.Println(mean, median)
	if mean == median {
		fmt.Println("simetris")
	} else if mean > median {
		fmt.Println("skew kanan")
	} else {
		fmt.Println("skew kiri")
	}
}