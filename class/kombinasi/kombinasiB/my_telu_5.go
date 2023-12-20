package main

import "fmt"

func main() {
	var poin, jumPoin, i, n int
	var gold, platinum, silver int
	var poinGold, poinPlatinum, poinSilver int
	var newUser bool
	var rataGold, rataPlatinum, rataSilver float64

	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		fmt.Scan(&poin, &newUser)
		for poin < 50 {
			fmt.Scan(&poin, &newUser)
		}
		if newUser {
			poin += 50
			if poin > 200 {
				poinGold += poin
				gold++
			} else if poin >= 100 {
				poinPlatinum += poin
				platinum++
			} else if poin >= 50 {
				poinSilver += poin
				silver++
			}
		} else {
			if poin > 200 {
				poinGold += poin
                gold++
			} else if poin >= 100 {
				poinPlatinum += poin
                platinum++
			} else if poin >= 50 {
				poinSilver += poin
                silver++
			}
		}
		jumPoin += poin
		if jumPoin > 500 {
			break
		}
		if gold == 0 {
			gold = 1
		} else if platinum == 0 {
			platinum = 1
        } else if silver == 0 {
			silver = 1
		}
		rataGold = float64(poinGold)/float64(gold)
		rataPlatinum = float64(poinPlatinum)/float64(platinum)
		rataSilver = float64(poinSilver)/float64(silver)
	}
	fmt.Println("gold: ", rataGold)
	fmt.Println("platinum: ", rataPlatinum)
	fmt.Println("silver: ", rataSilver)
}
