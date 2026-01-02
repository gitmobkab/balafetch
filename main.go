package main

import (
	"fmt"
	"math/rand/v2"
)

func pickRandomString(strs []string) string {
	var random_index int = rand.IntN(len(strs))
	return strs[random_index]
}

func pickRandomBalatroCardCategory() string{
	cards_categories := []string{
		"Jokers",
		"Tarot Cards",
		"Planet Cards",
		"Spectral Cards",
		"Vouchers",
	}

	return pickRandomString(cards_categories)
}


func main() {
	fmt.Println(pickRandomBalatroCardCategory())
}

