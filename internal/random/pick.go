package random

import (
	"math/rand"
	"github.com/gitmobkab/balafetch/internal/data"
)


type Picker struct{
	rng *rand.Rand
}

func NewPicker(seed int64) *Picker {
	return &Picker{
		rng: rand.New(rand.NewSource(seed)),
	}
}


/* 
PickInt returns a random integer in the range [0, max). it panics if max <= 0.
max is inclusive, so PickInt(1) will always return 0, PickInt(2) will return either 0 or 1, etc.

Naively calls rng.Intn(max), expected to be called in a 'safe' environnement*/
func (picker *Picker) PickInt(max int) int{
	return picker.rng.Intn(max)
}

// PickRandomString returns a random string from the given slice. it panics if the slice is empty.
// Naively calls PickInt(len(strs)) to get a random index, expected to be called in a 'safe' environnement
func (picker *Picker) PickRandomString(strs []string) string {
	var random_index int = picker.PickInt(len(strs))
	return strs[random_index]
}

// PickRandomBalatroCardCategory returns a random category from the list of balatro card categories.
// safest of the group, will only panic if the data.Cards_categories slice is empty.
func (picker *Picker) PickRandomBalatroCardCategory() string{
	return picker.PickRandomString(data.Cards_categories)
}
