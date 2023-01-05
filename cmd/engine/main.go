package main

import (
	"flag"
	"fmt"

	"simple-blockchain/node"
)

var configFlag = flag.String("config", "./config.json", "configuration json file path")

func main() {
	flag.Parse()
	config := NewConfig(*configFlag)

	fmt.Println("httpDialPort:", config.HttpDialPort)
	fmt.Println("httpListenPort:", config.HttpListenPort)
	nodeConfig := &node.Config{
		HttpListenPort: config.HttpListenPort,
		HttpDialPort:   config.HttpDialPort,
	}
	node, err := node.New(nodeConfig)
	if err != nil {
		panic(err)
	}
	node.Start()
}
