package main

import (
	"fmt"
	"github.com/gitmobkab/balafetch/internal/exit_codes"
)

func Help() int{
	fmt.Println("Balafetch - The stupid balatro flavoured fastfetch wrapper")
	fmt.Println("Usage: balafetch [-h | -v]")
	fmt.Println("Options:")
	fmt.Println("  -h      Show help information")
	fmt.Println("  -v 	   Show version information")
	return exitCodes.SuccessCode
}