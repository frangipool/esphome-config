//go:build ignore

// go generate
package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
	"text/template"
)

type templateData struct {
	Name         string
	Electrolyser bool
	Redox        bool
	PH           bool
	Booster      bool
}

//go:embed frangipool.yaml.tmpl
var templateConfig string

func main() {
	fmt.Println("generating frangipool configs...")
	versions := []string{"salt_booster_without_ph", "salt_without_ph", "salt_booster_without_ph_without_redox", "salt_without_ph_without_redox"}

	for _, version := range versions {
		data := templateData{}
		data.Name = version
		data.PH = true
		data.Redox = true

		if strings.Contains(version, "salt") {
			data.Electrolyser = true
		}
		if strings.Contains(version, "booster") {
			data.Booster = true
		}
		if strings.Contains(version, "without_ph") {
			data.PH = false
		}
		if strings.Contains(version, "without_redox") {
			data.Redox = false
		}

		tmpl, err := template.New("template").Parse(templateConfig)
		if err != nil {
			panic(err)
		}

		f, err := os.Create("frangipool_" + version + ".yaml")
		if err != nil {
			panic(err)
		}

		err = tmpl.Execute(f, data)
		if err != nil {
			panic(err)
		}
		err = f.Close()
		if err != nil {
			panic(err)
		}

	}

}
