package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	Port int
}

func NewConfig(file string) *Config {
	c := new(Config)

	if file, err := os.Open(file); err != nil {
		panic(err)
	} else {
		defer file.Close()
		if err := json.NewDecoder(file).Decode(c); err != nil {
			panic(err)
		} else {
			// c.sanitize()
			return c
		}
	}
}
