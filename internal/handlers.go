package internal

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"regexp"
	"strings"
	"time"
)

func (server *Server) ConnHanlder(conn net.Conn){
	defer conn.Close()
	Read(conn, "logo.txt")
	scanner := bufio.NewScanner(conn)
	username, err := server.TakeName(conn, scanner)
	if err != nil {
		fmt.Println("Error with handling username", err)
		return
	}
	if len(server.Users) > 9 {
		_, err := fmt.Fprintln(conn, "Chat is full, maximum 10 users allowed")
		if err != nil {
			log.Println(err)
			return
		}
		return
	}
	server.Lock()
	server.Users[conn] = username
	server.Unlock()

	Read(conn, "history.txt")
	_, err = fmt.Fprint(conn, fmt.Sprintf("[%s]:[%s]:", time.Now().Format("2006-01-02 15:04:05"), username))
	if err != nil {
		log.Println(err)
		return 
	}
	join := username + " has joined our chat"
	server.Msg <- Message{Text: join, Name: username}

	for scanner.Scan(){
		if len(strings.TrimSpace(scanner.Text())) == 0 || strings.TrimSpace(scanner.Text()) == "\n" || !regexp.MustCompile("^[\u0400-\u04FF\u0020-\u007F]+$").MatchString(scanner.Text()){
			_, err := fmt.Fprint(conn, "Please, text something\n")
			if err != nil {
				log.Fatal(err)
				return
			}
			_, err = fmt.Fprint(conn, fmt.Sprintf("[%s]:[%s]:", time.Now().Format("2006-01-02 15:04:05"), username))
			if err != nil {
				log.Fatal(err)
				return
			}
			continue
		} else {
			if scanner.Text() == "--changeName" {
				go server.NameChanger(username, conn)
				username = <-server.ChangeName
			} else {
				text := fmt.Sprintf("[%s]:[%s]:%s", time.Now().Format("2006-01-02 15:04:05"), username, scanner.Text())
				server.Msg <- Message{Text: text, Name: username}
				server.History.WriteString(text + "\n")
				_, err = fmt.Fprint(conn, fmt.Sprintf("[%s]:[%s]:", time.Now().Format("2006-01-02 15:04:05"), username))
				if err != nil {
					log.Fatal(err)
					return
				}

			}
		}
	}
	left := username + " has left our chat"
	server.Msg <- Message{Text: left, Name: username}
	server.Lock()
	delete(server.Users, conn)
	server.Unlock()
}

func (server *Server) TakeName(conn net.Conn, scanner *bufio.Scanner)(string, error){
	var username string
	server.Lock()
	for {
		_, err := fmt.Fprint(conn, "[ENTER YOUR NAME]:")
		if err != nil {
			return "", err
		}
		if scanner.Scan(){
			username = strings.TrimSpace(scanner.Text())
			alphanumeric := regexp.MustCompile(`^[a-zA-Z0-9]*$`)
			if !alphanumeric.MatchString(username){
				_, err := fmt.Fprintln(conn, "Please use only letters and numbers!")
				if err != nil {
					return "", err
				}
				continue
			}
			if !server.DuplicatedName(username){
				_, err := fmt.Fprintln(conn, "Name is already used, please try another one!")
				if err != nil {
					return "", err
				}
				continue
			}
			if username == ""{
				_, err := fmt.Fprintln(conn, "You didn't enter your name, please try again!")
				if err != nil {
					return "", err
				}
				continue
			}
			break
		}
	}
	server.Unlock()
	return username, nil
}

func (server *Server) DuplicatedName(name string) bool {
	for _, user := range server.Users {
		if user == name {
			return false
		}
	}
	return true
}


func (server *Server) NameChanger(oldname string, conn net.Conn) {
	s := bufio.NewScanner(conn)
	name, err := server.TakeName(conn, s)
	if err != nil {
		log.Fatal(err)
	}
	server.Lock()
	server.Users[conn] = name
	server.Unlock()
	server.History.WriteString(oldname + " has changed name to " + name + "\n")
	server.Msg <- Message{Name: oldname, Text: oldname + " has changed name to " + name}
	server.ChangeName <- name
}