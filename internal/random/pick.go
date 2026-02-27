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



func (picker *Picker) PickInt(max int) int{
	return picker.rng.Intn(max)
}

func (picker *Picker) PickRandomString(strs []string) string {
	var random_index int = picker.PickInt(len(strs))
	return strs[random_index]
}

func (picker *Picker) PickRandomBalatroCardCategory() string{
	return picker.PickRandomString(data.Cards_categories)
}
