package combinations

func Combinations(yollar [][]string, indeks int, secili []int, maxYol int, enIyiKombinasyon []int) ([]int, int) {
	if len(secili) > maxYol {
		maxYol = len(secili)
		enIyiKombinasyon = make([]int, len(secili))
		copy(enIyiKombinasyon, secili) // Dilimin korunmasını sağlamak için.
	}

	for i := indeks; i < len(yollar); i++ {
		cakisiyor := false
		for _, s := range secili {
			if YollarCakisiyor(yollar[s], yollar[i]) {
				cakisiyor = true
				break
			}
		}
		if !cakisiyor {
			secili = append(secili, i)
			enIyiKombinasyon, maxYol = Combinations(yollar, i+1, secili, maxYol, enIyiKombinasyon)
			secili = secili[:len(secili)-1] // Mevcut yol seçilmiş yollardan çıkarılır ve diğer kombinasyonlar için yeni aramalar yapılır.
		}
	}
	return enIyiKombinasyon, maxYol
}

func YollarCakisiyor(yol1, yol2 []string) bool {
	kume := make(map[string]bool)
	for _, oda := range yol1[1 : len(yol1)-1] { // Başlangıç ve bitişi hariç tut
		kume[oda] = true
	}
	for _, oda := range yol2[1 : len(yol2)-1] {
		if kume[oda] {
			return true
		}
	}
	return false
}
