package service

import (
	"IM/logic"
	"encoding/json"
	"fmt"
	"sync"
)

var MyServer = NewServer()

type Server struct {
	Clients   map[uint64]*Client
	mutex     *sync.Mutex
	Broadcast chan []byte
	Register  chan *Client
	Ungister  chan *Client
}

func NewServer() *Server {
	return &Server{
		mutex:     &sync.Mutex{},
		Clients:   make(map[uint64]*Client),
		Broadcast: make(chan []byte),
		Register:  make(chan *Client, 10240),
		Ungister:  make(chan *Client),
	}
}

func (s *Server) Start() {
	fmt.Println("start server")
	for {
		select {
		case conn := <-s.Register:
			fmt.Println("server.go line:35,login, ", conn.Uid)
			s.Clients[conn.Uid] = conn
			msg := &logic.Message{
				From:    0,
				To:      []uint64{conn.Uid},
				Content: "welcome!",
			}
			data, _ := json.Marshal(msg)
			conn.Send <- data
		case conn := <-s.Ungister:
			fmt.Println("loginout", conn.Uid)
			if _, ok := s.Clients[conn.Uid]; ok {
				close(conn.Send)
				delete(s.Clients, conn.Uid)
			}
		}
	}
}

func (s *Server) GetConn(uid uint64) *Client {
	return s.Clients[uid]
}
