package main

import (
	"fmt"
	"github.com/gitmobkab/balafetch/internal/data"
)

func Version() {
	fmt.Printf("%s %s\n", data.AppName, data.GetVersion()) 
}