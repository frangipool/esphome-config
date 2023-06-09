//go:build ignore

// go generate
package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"text/template"

	"gopkg.in/yaml.v3"
)

type devices struct {
	Device []deviceConfig
}

type deviceConfig struct {
	Name         string `yaml:"name"`
	Electrolyser bool   `yaml:"electrolyser"`
	Redox        bool   `yaml:"redox"`
	PH           bool   `yaml:"ph"`
	Booster      bool   `yaml:"booster"`
}

//go:embed frangipool.yaml.tmpl
var templateConfig string

//go:embed config.yaml
var configsYaml string

//go:embed README.md.tmpl
var templateReadme string

func main() {
	fmt.Println("generating frangipool configs...")

	var d devices
	// Parse yaml config
	if err := yaml.Unmarshal([]byte(configsYaml), &d.Device); err != nil {
		log.Fatalf("error: %v", err)
	}

	// Generate esphome config files
	for _, device := range d.Device {

		tmpl, err := template.New("template").Parse(templateConfig)
		if err != nil {
			panic(err)
		}

		f, err := os.Create("frangipool_" + device.Name + ".yaml")
		if err != nil {
			panic(err)
		}

		err = tmpl.Execute(f, device)
		if err != nil {
			panic(err)
		}
		err = f.Close()
		if err != nil {
			panic(err)
		}

	}

	// Generate README
	tmpl, err := template.New("template").Parse(templateReadme)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("README.md")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(f, d)
	if err != nil {
		panic(err)
	}
	err = f.Close()
	if err != nil {
		panic(err)
	}
}
