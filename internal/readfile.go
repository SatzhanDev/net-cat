package internal

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func Read(conn net.Conn, str string) {
	file, err := os.Open(str)
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(file)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			break
		}
		fmt.Fprint(conn, line)
	}
}
