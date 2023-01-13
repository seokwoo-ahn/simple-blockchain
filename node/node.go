package node

import (
	"fmt"
	"simple-blockchain/p2p"
	"sync"
)

type Node struct {
	httpListenPort int
	httpDialPort   int
}

func New(config *Config) (*Node, error) {
	if config.HttpListenPort == int(0) || config.HttpDialPort == int(0) {
		return &Node{}, fmt.Errorf("need http port num")
	}
	node := &Node{
		httpListenPort: config.HttpListenPort,
		httpDialPort:   config.HttpDialPort,
	}
	return node, nil
}

func (n *Node) Start() {
	var wg sync.WaitGroup
	wg.Add(2)

	p2pConfig := &p2p.Config{
		HttpListenPort: n.httpListenPort,
		HttpDialPort:   n.httpDialPort,
	}

	server, _ := p2p.NewServer(p2pConfig)
	client, _ := p2p.NewClient(p2pConfig)
	client.Wg = &wg

	go server.Start()
	go client.Start()

	wg.Wait()
}
