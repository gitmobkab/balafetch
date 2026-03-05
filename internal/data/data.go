package data

import (
	"runtime"
	"fmt"
	"strings"
	"sort"
)

const AppName string = "balafetch"

const OS string = runtime.GOOS
const ARCH string = runtime.GOARCH

var Cards_categories = []string{
	"jokers",
	"tarot cards",
	"planet cards",
	"spectral cards",
	"vouchers",
}

var CategoryResolution = map[string]string{
    // canonical → canonical
    "jokers":          "jokers",
    "tarot cards":     "tarot cards",
    "planet cards":    "planet cards",
    "spectral cards":  "spectral cards",
    "vouchers":        "vouchers",
    
	// aliases → canonical
    "joker":           "jokers",

    "tarots":          "tarot cards",
    "tarot":           "tarot cards",
    
	"planets":         "planet cards",
	"planet":          "planet cards",
	
	"spectrals":       "spectral cards",
    "spectral":        "spectral cards",
    
	"voucher":         "vouchers",
}

var CategoryAliases = map[string][]string{}

// Init populates the CategoryAliases map based on the CategoryResolution map.
// Init should be called before using CategoryAliases, and it is also called in DisplayCategoryHelp to ensure it's populated when displaying help.
// This allows us to easily display aliases for each category when showing help.
func Init() {
    for alias, canonical := range CategoryResolution {
        if alias != canonical {
            CategoryAliases[canonical] = append(CategoryAliases[canonical], alias)
			sort.Strings(CategoryAliases[canonical]) // to fix random order of aliases
        }
    }
}

// DisplayCategoryHelp is a helper function that displays the available categories and their aliases
// It avoids the need to manually init CategoryAliases and handle the logic of displaying aliases in multiple places.
func DisplayCategoryHelp() {
	Init()
	fmt.Println("Available categories (aliases in parentheses):")
		for _, category := range Cards_categories {
		aliases := CategoryAliases[category]
		if len(aliases) > 0 {
			fmt.Printf(" - %s (%s)\n", category, strings.Join(aliases, ", "))
		} else {
			fmt.Printf(" - %s\n", category)
		}
	}
}