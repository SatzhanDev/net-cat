# Net-cat

## Description

This project consists on recreating the NetCat in a Server-Client Architecture that can run in a server mode on a specified port listening for incoming connections, and it can be used in client mode, trying to connect to a specified port and transmitting information to the server.
NetCat, nc system command, is a command-line utility that reads and writes data across network connections using TCP or UDP. It is used for anything involving TCP, UDP, or UNIX-domain sockets, it is able to open TCP connections, send UDP packages, listen on arbitrary TCP and UDP ports and many more.

## Links

- [About net-cat](https://www.commandlinux.com/man-page/man1/nc.1.html)
- [How to write a Good readme](https://readme.so/editor)

## Author:

- @skanatba
- @dyesselb

## Usage:

Server:

      "go run ./cmd/main" or "go run ./cmd/main 4444"
      Listening on the port :8989

Client:

      "nc localhost 8989" or "nc localhost [another port number]"
