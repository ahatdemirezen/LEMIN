package filterpathways

import "lemin/combinations"

func FilterPathWays(yollar [][]string, karincaSayisi int) [][]string {
	var filtrelenmisYollar [][]string
	maxYol := 0
	var enIyiKombinasyon []int

	enIyiKombinasyon, _ = combinations.Combinations(yollar, 0, []int{}, maxYol, enIyiKombinasyon)

	for _, indeks := range enIyiKombinasyon {
		filtrelenmisYollar = append(filtrelenmisYollar, yollar[indeks])
		if len(filtrelenmisYollar) == karincaSayisi {
			break
		}
	}

	return filtrelenmisYollar
}
