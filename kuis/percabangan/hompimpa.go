package main
import "fmt"

func main()  {
	var hom, pim, pa string
	var pemenang string

	fmt.Scan(&hom, &pim, &pa)
	if (hom == pim && pim == pa) {
		pemenang = "imbang"
	} else if hom != pim && pim == pa {
		pemenang = "pemain 1 pemenang"
	} else if hom == pa && hom != pim && pim != pa{
		pemenang = "pemain 2 pemenang"
	} else if hom == pim && pim != pa {
		pemenang = "pemain 3 pemenang"
	} 
	fmt.Println(pemenang)
}