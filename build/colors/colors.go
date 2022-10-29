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

func MochaColor(flag string) {
	c := Color{
		Name:      "Catppuccin Mocha",
		Shortname: "mocha",
		Dark:      true,
		Rosewater: "#f5e0dc",
		Flamingo:  "#f2cdcd",
		Pink:      "#f5c2e7",
		Mauve:     "#cba6f7",
		Red:       "#f38ba8",
		Maroon:    "#eba0ac",
		Yellow:    "#f9e2af",
		Peach:     "#fab387",
		Green:     "#a63a1",
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

func MacchiatoColor(flag string) {
	c := Color{
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

func initTemplate(colorStruct Color, flag string) (string, error) {

	currentDir, _ := os.Getwd()
	templatePath := path.Join(currentDir, "build")
	themeFileName := flag + "_desktop"
	fileTemp, err := os.OpenFile(path.Join(templatePath, "template.go.tpl"), os.O_RDONLY, 0o644)
	if err != nil {
		panic(err)
	}
	tmpl, _ := template.ParseFiles(fileTemp.Name())
	themeLoc := path.Join(currentDir, "src", themeFileName)
	f, _ := os.Create(themeLoc)

	template := colorStruct
	err = tmpl.Execute(f, template)
	if err != nil {
		panic(err)
	}
	f.Close()
	returnString := "Mocha scheme generated."
	return returnString, nil
	// This is not complete please no judge me :kekw:
}
