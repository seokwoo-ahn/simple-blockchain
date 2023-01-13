package p2p

import (
	"fmt"
	"io"
	"net"
	"strconv"
)

type Server struct {
	httpListenPort int
}

func NewServer(config *Config) (*Server, error) {
	if config.HttpListenPort == int(0) {
		return &Server{}, fmt.Errorf("need http port num")
	}
	server := &Server{
		httpListenPort: config.HttpListenPort,
	}
	return server, nil
}

func (srv *Server) Start() {
	listenPort := ":" + strconv.Itoa(srv.httpListenPort)
	var l net.Listener

	l, err := net.Listen("tcp", listenPort)
	if nil != err {
		panic(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			continue
		}
		ConnHandler(conn)
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
