package internal

import (
	"fmt"
	"log"
	"time"
)


func (server *Server) Broadcaster(){
	for {
		select {
		case msg := <-server.Msg:
			server.Lock()
			for conn, name := range server.Users {
				if name != msg.Name {
					_, err := fmt.Fprintln(conn, "\n"+msg.Text)
					if err != nil {
						log.Println(err)
					}
				} else {
					continue
				}
				_, err := fmt.Fprint(conn, fmt.Sprintf("[%s]:[%s]:", time.Now().Format("2006-01-02 15:04:05"), name))
				if err != nil {
					log.Println(err)
				}
			}
			server.Unlock()
		}
	}
}