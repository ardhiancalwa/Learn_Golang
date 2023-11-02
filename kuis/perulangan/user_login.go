package main
import "fmt"

func main()  {
	var n int
	var username, password string

	n = 0
	for {
		fmt.Scan(&username, &password)
		if username == "admin" && password == "admin" {
			break
		}
		n++
	}
	fmt.Println(n, "Login berhasil")
}