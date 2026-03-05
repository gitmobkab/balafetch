package main

import (
	"fmt"
	"github.com/gitmobkab/balafetch/internal/data"
)

func Version(detailed bool) {
	if detailed {
		fmt.Printf("%s %s\n", data.AppName, data.GetDetailedVersion())
		return
	}
	fmt.Printf("%s %s\n", data.AppName, data.GetVersion()) 
}