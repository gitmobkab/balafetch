package data

import (
	"runtime"
	"fmt"
	"strings"
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

func Init() {
    for alias, canonical := range CategoryResolution {
        if alias != canonical {
            CategoryAliases[canonical] = append(CategoryAliases[canonical], alias)
        }
    }
}

func DisplayCategoryHelp() {
	Init()
	fmt.Println("Available categories (aliases in parantheses):")
		for _, category := range Cards_categories {
		aliases := CategoryAliases[category]
		if len(aliases) > 0 {
			fmt.Printf(" - %s  (%s)\n", category, strings.Join(aliases, ", "))
		} else {
			fmt.Printf(" - %s\n", category)
		}
	}
}