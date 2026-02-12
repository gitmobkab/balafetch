package main

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"github.com/gitmobkab/balafetch/internal/data"
)

func Version() {
	myFigure := figure.NewColorFigure(data.AppName, "cybermedium", "blue", true)
  	myFigure.Print()
	fmt.Println("The stupid balatro flavoured fastfetch wrapper")
	fmt.Printf("version: %s\n", data.VERSION)
	fmt.Printf("Under the %s License\n", data.LICENSE)
}