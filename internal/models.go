package internal

import (
	"net"
	"os"
	"sync"
)
type Message struct {
	Text string
	Name string
}

type Server struct{
	sync.Mutex
	Users      map[net.Conn]string
	Msg        chan Message
	History    *os.File
	ChangeName chan string
}