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
	data.DisplayCategoryHelp()
}