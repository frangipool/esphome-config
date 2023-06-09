package main

import (
	// Ensure all dependencies are imported for go generate.
	_ "gopkg.in/yaml.v3"
)

//go:generate go run template/main.go
