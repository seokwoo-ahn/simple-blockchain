package p2p

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

type Client struct {
	HttpDialPort int
	Wg           *sync.WaitGroup
}

func NewClient(config *Config) (*Client, error) {
	if config.HttpDialPort == int(0) {
		return &Client{}, fmt.Errorf("need http port num")
	}
	client := &Client{
		HttpDialPort: config.HttpDialPort,
	}
	return client, nil
}

func (c *Client) Start() {
	dialPort := ":" + strconv.Itoa(c.HttpDialPort)

	var d net.Conn
	var err error

	for {
		d, err = net.Dial("tcp", dialPort)
		if nil == err {
			fmt.Println("connect complete!!")
			break
		}
		time.Sleep(3 * time.Second)
		fmt.Println("try dial")
	}

	d.Write([]byte("hello"))
	c.Wg.Done()
}
