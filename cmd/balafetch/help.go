package main

import (
	"fmt"
	pflag "github.com/spf13/pflag"
)

func Help() {
	fmt.Println("Balafetch - The stupid balatro flavoured fastfetch wrapper")
	fmt.Println("Usage: balafetch [options]")
	fmt.Println("Options:")
	pflag.PrintDefaults()
}