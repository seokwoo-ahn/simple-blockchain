package main

import (
	"flag"
	"fmt"
)

var configFlag = flag.String("config", "./config.json", "configuration json file path")

func main() {
	flag.Parse()
	config := NewConfig(*configFlag)
	fmt.Println(config.Port)
}
