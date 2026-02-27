package main

import (
	"fmt"
	pflag "github.com/spf13/pflag"
	"github.com/gitmobkab/balafetch/internal/data"
)

func Help() {
	fmt.Printf("%s - The stupid balatro flavoured fastfetch wrapper\n",data.AppName)
	fmt.Printf("Usage: %s [options] [card_category]\n",data.AppName)
	fmt.Println("Options:")
	pflag.PrintDefaults()
	fmt.Println("Card Categories:")
	for _, category := range data.Cards_categories {
		fmt.Printf(" - %s\n", category)
	}
	fmt.Println("\nNote: use double quotes for categories with spaces, e.g. \"tarot cards\"")
}