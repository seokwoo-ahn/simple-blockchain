package node

import (
	"fmt"
	"io"
	"net"
	"strconv"
	"time"
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
	listenPort := ":" + strconv.Itoa(n.httpListenPort)
	dialPort := ":" + strconv.Itoa(n.httpDialPort)

	var l net.Listener
	var d net.Conn

	l, err := net.Listen("tcp", listenPort)
	if nil != err {
		panic(err)
	}
	for {
		d, err = net.Dial("tcp", dialPort)
		if nil == err {
			fmt.Println("connect complete!!")
			break
		}
		time.Sleep(3 * time.Second)
		fmt.Println("try dial")
	}

	var conn net.Conn
	defer conn.Close()
	for {
		conn, err = l.Accept()
		if err != nil {
			continue
		}
		d.Write([]byte("hello!!"))
		go ConnHandler(conn)
	}
}

func ConnHandler(conn net.Conn) {
	recvBuf := make([]byte, 4096)
	for {
		n, err := conn.Read(recvBuf)
		if nil != err {
			if io.EOF == err {
				return
			}
			return
		}
		if 0 < n {
			data := recvBuf[:n]
			fmt.Println("data received!!", data)
		}
	}
}
