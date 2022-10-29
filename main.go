package main

import (
	"flag"

	"github.com/catppuccin/telegram/build/colors"
)

func main() {
	var colorFlag string
	flag.StringVar(&colorFlag, "color", "mocha", "some value")
	flag.Parse()
	if colorFlag == "mocha" {
		colors.MochaColor(colorFlag)
	} else if colorFlag == "macchiato" {
		colors.MacchiatoColor(colorFlag)
	} else if colorFlag == "frappe" {
		colors.FrappeColor(colorFlag)
	} else {
		colors.LatteColor(colorFlag)
	}
}
