package simulateAntMovement

import (
	"fmt"
	"strings"
)

func SimulateAntMovement(yollar [][]string, karincaSayisi int, başlangiç, hedef string, enKisaYol []string) []string {
	var hareketler []string
	karincaPozisyonu := make(map[int]int)
	karincaHedefte := make(map[int]bool)
	karincaYollari := make(map[int][]string)
	aktifKarincaSayisi := karincaSayisi

	for i := 1; i <= karincaSayisi; i++ {
		if i == karincaSayisi {
			karincaYollari[i] = enKisaYol // Son karınca en kısa yolu takip eder
		} else {
			karincaYollari[i] = yollar[(i-1)%len(yollar)] // diğerleri sırayla takip eder.
		}
		karincaPozisyonu[i] = 0   // bütün karıncalar başlangıç pozisyonundadır.
		karincaHedefte[i] = false // hiçbir karınca bitişe ulaşmamıştır.
	}

	tur := 0
	for aktifKarincaSayisi > 0 {
		tur++
		var turHareketi []string
		tünelKullanimi := make(map[string]bool)

		for i := 1; i <= karincaSayisi; i++ {
			if karincaHedefte[i] {
				continue
			}

			şuAnkiOda := karincaYollari[i][karincaPozisyonu[i]]
			sonrakiOda := karincaYollari[i][karincaPozisyonu[i]+1]
			tünel := fmt.Sprintf(şuAnkiOda, sonrakiOda)
			tersTünel := fmt.Sprintf(sonrakiOda, şuAnkiOda)

			if !tünelKullanimi[tünel] && !tünelKullanimi[tersTünel] {
				turHareketi = append(turHareketi, fmt.Sprintf("L%d-%s", i, sonrakiOda))
				tünelKullanimi[tünel] = true
				tünelKullanimi[tersTünel] = true
				karincaPozisyonu[i]++
				if sonrakiOda == hedef {
					karincaHedefte[i] = true
					aktifKarincaSayisi--
				}
			}
		}

		if len(turHareketi) > 0 {
			hareketler = append(hareketler, strings.Join(turHareketi, " "))
		} else {
			break // Eğer bu turda hareket eden karınca yoksa, döngüyü kır
		}
	}
	return hareketler
}
