package bonus

//Joãozinho ganhou um kit de construção de torres. O kit é composto por várias barras de madeira, e seus comprimentos são
//conhecidos. As barras podem ser empilhadas umas sobre as outras, desde que seus comprimentos sejam iguais.
//
//Joãozinho quer construir o menor número possível de torres com as barras que tem. Você deve ajudar Joãozinho a usar as
//barras da melhor maneira possível, determinando a altura da torre mais alta e quantas torres podem ser construídas.

func CalculateTowers(barLengths []int) (int, int) {
	for i := 0; i < len(barLengths); i++ {
		for j := 0; j < len(barLengths); j++ {
			if barLengths[i] > barLengths[j] {
				barLengths[i], barLengths[j] = barLengths[j], barLengths[i]
			}
		}
	}

	var (
		maxTowerHeight     int
		towersCount        = 1
		currentTowerHeight = 1
	)
	for i := 1; i < len(barLengths); i++ {
		if barLengths[i] == barLengths[i-1] {
			currentTowerHeight++
		} else {
			if currentTowerHeight > maxTowerHeight {
				maxTowerHeight = currentTowerHeight
			}
			currentTowerHeight = 1
			towersCount++
		}
	}

	if currentTowerHeight > maxTowerHeight {
		maxTowerHeight = currentTowerHeight
	}

	return maxTowerHeight, towersCount
}
