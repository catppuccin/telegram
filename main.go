package main

import (
	"flag"
	"fmt"

	"github.com/catppuccin/telegram/build/colors"
)

func main() {
	var colorFlag string
	flag.StringVar(&colorFlag, "color", "mocha", "some value")
	var accentFlag string
	flag.StringVar(&accentFlag, "accent", "ctpGreen", "some value")
	flag.Parse()
	if colorFlag == "mocha" {
		colors.MochaColor(colorFlag, accentFlag)
	} else if colorFlag == "macchiato" {
		colors.MacchiatoColor(colorFlag, accentFlag)
	} else if colorFlag == "frappe" {
		colors.FrappeColor(colorFlag, accentFlag)
	} else if colorFlag == "latte" {
		colors.LatteColor(colorFlag, accentFlag)
	} else {
		fmt.Println("You've entered the wrong color.")
	}
}
