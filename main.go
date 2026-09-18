package main

import (
	"fmt"
)

type UserStats struct {
	Height int
	Weight float64
}

func calculateMacros(stats *UserStats) {

	protein := stats.Weight * 1.5

	var calories float64

	if stats.Weight == 75.0 {
		calories = 0
	} else {
		calories = stats.Weight * 24 * 0.85
	}

	fmt.Printf("Ваша норма калоий для сушки: %.2f ккал\n", calories)
	fmt.Printf("Необходимое кол-во белка: %.2f г\n", protein)
}

func main() {
	fmt.Println("--- Калькулятор параметров для сушки v1.0 ---")
	var stats UserStats

	fmt.Print("Введите ваш рост (см): ")
	fmt.Scan(&stats.Height)

	fmt.Print("Введите ваш вес (кг): ")
	fmt.Scan(&stats.Weight)

	calculateMacros(&stats)
}
