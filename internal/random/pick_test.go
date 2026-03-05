package random

import (
	"testing"
	"github.com/gitmobkab/balafetch/internal/data"
)

func TestPickInt_AlwaysInBounds(t *testing.T) {
	picker := NewPicker(42)
	max := 10
	for range 40 {
		result := picker.PickInt(max)
		if result < 0 || result >= max {
			t.Errorf("PickInt(%d) = %d, out of bounds [0, %d)", max, result, max)
		}
	}
}

func TestPickInt_IsDeterministicWithSameSeed(t *testing.T) {
	p1 := NewPicker(99)
	p2 := NewPicker(99)
	for range 20 {
		if p1.PickInt(100) != p2.PickInt(100) {
			t.Error("same seed produced different results")
		}
	}
}

func TestPickRandomString_AlwaysReturnsValidElement(t *testing.T) {
	strs := []string{"alpha", "beta", "gamma", "delta"}
	strSet := map[string]bool{}
	for _, s := range strs {
		strSet[s] = true
	}

	picker := NewPicker(7)
	for range 50 {
		result := picker.PickRandomString(strs)
		if !strSet[result] {
			t.Errorf("PickRandomString returned unexpected value %q", result)
		}
	}
}

func TestPickRandomBalatroCardCategory_ReturnsValidCategory(t *testing.T) {
	validCategories := map[string]bool{}
	for _, c := range data.Cards_categories {
		validCategories[c] = true
	}

	picker := NewPicker(0)
	for range 50 {
		result := picker.PickRandomBalatroCardCategory()
		if !validCategories[result] {
			t.Errorf("PickRandomBalatroCardCategory returned unknown category %q", result)
		}
	}
}