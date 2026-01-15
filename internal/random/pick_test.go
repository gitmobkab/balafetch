package random_test

import (
	"testing"
	"github.com/gitmobkab/balafetch/internal/random"
	"slices"
)

func TestPickIntDeterministic(tester *testing.T){
	const SEED int64 = 5
	const TARGET int = 26
	const PICK_RANGE int = 50
	picker := random.NewPicker(SEED)
	if out := picker.PickInt(PICK_RANGE); out != TARGET{
		tester.Fatalf("Expected %d, got %d for SEED: %d with RANGE: %d",
						TARGET, out, SEED, PICK_RANGE)
	}
}

func TestPickIntRange(tester *testing.T){
	const SEED int64 = 1 
	const MAX int = 35
	picker := random.NewPicker(SEED)
	for range 1000{
		out := picker.PickInt(MAX)
		if out < 0 || out > MAX {
			tester.Fatalf("Out of range, got %d for SEED: %d with MAX: %d", 
							out, SEED, MAX)
		}
	}
}

func TestPickRandomStringDeterministic(tester *testing.T){
	const SEED int64 = 42
	const EXPECTED string = "banana"
	testStrings := []string{"apple", "banana", "cherry", "date"}
	picker := random.NewPicker(SEED)
	result := picker.PickRandomString(testStrings)
	
	if result != EXPECTED{
		tester.Fatalf("Unexpected result, got %s for SEED: %d \nExpected %s", result, SEED, EXPECTED)
	}
}

func TestPickRandomStringRange(tester *testing.T){
	const SEED int64 = 10
	testStrings := []string{"red", "green", "blue", "yellow", "orange"}
	picker := random.NewPicker(SEED)
	
	for range 1000 {
		result := picker.PickRandomString(testStrings)
		if !slices.Contains(testStrings, result){
			tester.Fatalf("Expected valid choice, got %s with SEED: %d \noptions: %#v",result, SEED, testStrings)
		}
	}
}

func TestPickRandomBalatroCardCategoryDeterministic(tester *testing.T){
	const SEED int64 = 7
	const EXPECTED string = "Tarot Cards"
	picker := random.NewPicker(SEED)
	result := picker.PickRandomBalatroCardCategory()
	
	validCategories := []string{"Jokers", "Tarot Cards", "Planet Cards", "Spectral Cards", "Vouchers"}
	
	if result != EXPECTED {
		tester.Fatalf("Expected valid card category, got %s for SEED: %d\nValids: %#v", result, SEED, validCategories)
	}
}



