package main
import "fmt"

func main()  {
    var i, n, keuntungan, pendapatan, aset, bagiKeuntungan int

    fmt.Scan(&n)
    i = 1
    aset = 0
    for i <= n {
        fmt.Scan(&keuntungan)
        bagiKeuntungan = keuntungan / 3
        pendapatan += bagiKeuntungan
        aset += keuntungan % 3
        i++
    }
    fmt.Println(pendapatan, aset)
}
