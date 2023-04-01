package internal

import (
	"fmt"
	"log"
	"net"
	"os"
)

func (server *Server)RunServer(port string){
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Println(err)
		return
	}
	defer listener.Close()
	server.History, err = os.Create("history.txt")
	if err != nil {
		log.Println(err)
		return
	}
	go server.Broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go server.ConnHanlder(conn)
	}
}
