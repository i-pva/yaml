package main

import (
	"fmt"
	"reflect"

	v3 "github.com/i-pva/yaml"
)

type d struct {
	Fields  map[string]string `yaml:"fields"`
	Fieldds map[string]string `yaml:"fieldds"`
}
type Config struct {
	ss       *d
	Fields   map[string]string `yaml:"fields"`
	Fddields map[string]string `yaml:"fdields"`
}

const configYAML = `
fields:
        f1: null
        f2: 
`

func main() {
	cfg := Config{}

	fmt.Println(reflect.TypeOf(cfg).Size())

	err := v3.Unmarshal([]byte(configYAML), &cfg)
	if err != nil {
		panic(err)
	}

	fmt.Println("v3", cfg.Fields)
}
