package colors

import (
	"os"
	"path"
	"text/template"
)

type Color struct {
	Name      string
	Shortname string
	Dark      bool
	Wallpaper string
	Accent    string
	Rosewater string
	Flamingo  string
	Pink      string
	Mauve     string
	Red       string
	Maroon    string
	Peach     string
	Yellow    string
	Green     string
	Teal      string
	Sky       string
	Sapphire  string
	Blue      string
	Lavender  string
	Text      string
	Subtext0  string
	Subtext1  string
	Overlay2  string
	Overlay1  string
	Overlay0  string
	Surface2  string
	Surface1  string
	Surface0  string
	Base      string
	Mantle    string
	Crust     string
}

func MochaColor(flag string, accent string) {
	c := Color{
		Name:      "Catppuccin Mocha",
		Shortname: "mocha",
		Dark:      true,
		Wallpaper: "t.me/bg/iZumGdCBBVEBAAAAkZ7Gl0zmXZk",
		Accent:    accent,
		Rosewater: "#f5e0dc",
		Flamingo:  "#f2cdcd",
		Pink:      "#f5c2e7",
		Mauve:     "#cba6f7",
		Red:       "#f38ba8",
		Maroon:    "#eba0ac",
		Yellow:    "#f9e2af",
		Peach:     "#fab387",
		Green:     "#a6e3a1",
		Teal:      "#94e2d5",
		Sky:       "#89dceb",
		Sapphire:  "#74c7e6",
		Blue:      "#89b4fa",
		Lavender:  "#b4befe",
		Text:      "#cdd6f4",
		Subtext1:  "#bac2de",
		Subtext0:  "#a6adc8",
		Overlay2:  "#9399b2",
		Overlay1:  "#7f4849c",
		Overlay0:  "#6c7086",
		Surface2:  "#585b70",
		Surface1:  "#45475a",
		Surface0:  "#313244",
		Base:      "#1e1e2e",
		Mantle:    "#181825",
		Crust:     "#11111b",
	}
	initTemplate(c, flag)
}

func MacchiatoColor(flag string, accent string) {
	c := Color{
		Name:      "Catpppuccin Macchiato",
		Shortname: "macchiato",
		Dark:      true,
		Wallpaper: "t.me/bg/wrjDzHXw_VADAAAAYTt1ohXDJN4",
		Accent:    accent,
		Rosewater: "#f5e0dc",
		Flamingo:  "#f0c6c6",
		Pink:      "#f5c2e7",
		Mauve:     "#c6a0f6",
		Red:       "#ed8796",
		Maroon:    "#ee99a0",
		Peach:     "#f5a97f",
		Yellow:    "#eed49f",
		Green:     "#a6da95",
		Teal:      "#8bd5ca",
		Sky:       "#91d7e3",
		Sapphire:  "#7dc4e4",
		Blue:      "#8aadf4",
		Lavender:  "#b7bdf8",
		Text:      "#cad3f5",
		Subtext1:  "#b8c0e0",
		Subtext0:  "#a5adcb",
		Overlay2:  "#939ab7",
		Overlay1:  "#8087a2",
		Overlay0:  "#6e738d",
		Surface2:  "#5b6078",
		Surface1:  "#494d64",
		Surface0:  "#363a4f",
		Base:      "#24273a",
		Mantle:    "#1e2030",
		Crust:     "#181926",
	}
	initTemplate(c, flag)
}

func FrappeColor(flag string, accent string) {
	c := Color{
		Name:      "Catppuccin Frappe",
		Shortname: "frappe",
		Dark:      true,
		Wallpaper: "t.me/bg/psvRMp-j_VACAAAA6D6zUK3aWhs",
		Accent:    accent,
		Rosewater: "#f2d5cf",
		Flamingo:  "#eebebe",
		Pink:      "#f4b8e4",
		Mauve:     "#ca9ee6",
		Red:       "#e78284",
		Maroon:    "#ea999c",
		Peach:     "#ef9f76",
		Yellow:    "#e5c890",
		Green:     "#a6d189",
		Teal:      "#81c8be",
		Sky:       "#99d1db",
		Sapphire:  "#85c1dc",
		Blue:      "#8caaee",
		Lavender:  "#babbf1",
		Text:      "#c6d0f5",
		Subtext1:  "#b5bfe2",
		Subtext0:  "#a5adce",
		Overlay2:  "#949cbb",
		Overlay1:  "#838ba7",
		Overlay0:  "#737994",
		Surface2:  "#626880",
		Surface1:  "#51576d",
		Surface0:  "#414559",
		Base:      "#303446",
		Mantle:    "#292c3c",
		Crust:     "#232634",
	}
	initTemplate(c, flag)
}

func LatteColor(flag string, accent string) {
	c := Color{
		Name:      "Catppuccin Latte",
		Shortname: "latte",
		Dark:      false,
		Wallpaper: "t.me/bg/StNc4pYj_FABAAAAOsWwdh_6J6I",
		Accent:     accent,
		Rosewater: "#dc8a78",
		Flamingo:  "#dd7878",
		Pink:      "#ea76cb",
		Mauve:     "#8839ef",
		Red:       "#d20f39",
		Maroon:    "#e64553",
		Peach:     "#fe640b",
		Yellow:    "#df8e1d",
		Green:     "#40a02b",
		Teal:      "#179299",
		Sky:       "#04a5e5",
		Sapphire:  "#209fb5",
		Blue:      "#1e66f5",
		Lavender:  "#7287fd",
		Text:      "#4c4f69",
		Subtext1:  "#5c5f77",
		Subtext0:  "#6c6f85",
		Overlay2:  "#7c7f93",
		Overlay1:  "#8c8fa1",
		Overlay0:  "#9ca0b0",
		Surface2:  "#acb0be",
		Surface1:  "#bcc0cc",
		Surface0:  "#ccd0da",
		Base:      "#eff1f5",
		Mantle:    "#e6e9ef",
		Crust:     "#dce0e8",
	}
	initTemplate(c, flag)
}
func initTemplate(colorStruct Color, flag string) {

	currentDir, _ := os.Getwd()
	templatePath := path.Join(currentDir, "build")
	fileTemp, err := os.OpenFile(path.Join(templatePath, "template.go.tpl"), os.O_RDONLY, 0o644)
	if err != nil {
		panic(err)
	}
	tmpl, _ := template.ParseFiles(fileTemp.Name())
	themeFile := "desktop"
	themeLoc := path.Join(currentDir, "src", flag, themeFile)
	f, _ := os.Create(themeLoc)

	template := colorStruct
	err = tmpl.Execute(f, template)
	if err != nil {
		panic(err)
	}
	f.Close()
}
