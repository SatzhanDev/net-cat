package main

import (
	"fmt"
	"net"
	"net-cat/internal"
	"net-cat/pkg"
	"os"
)



func main(){
	port := ""
	if len(os.Args) == 1 {
		port = ":8989"
	} else if len(os.Args) == 2 && pkg.Chekport(os.Args[1]) {
		port = ":" + os.Args[1]
	} else {
		fmt.Println("[USAGE]: go run . $port")
		return
	}
	server := &internal.Server{
		Users:      make(map[net.Conn]string),
		Msg:        make(chan internal.Message),
		History:    &os.File{},
		ChangeName: make(chan string),
	}
	fmt.Println("Listening on the port", port)
	server.RunServer(port)
}