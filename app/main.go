package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

var _ = net.Listen
var _ = os.Exit

func main() {
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	req := make([]byte, 1024)
	conn.Read(req)

	lines := strings.Split(string(req), "\r\n")
	method := strings.Split(lines[0], " ")[0]
	path := strings.Split(lines[0], " ")[1]
	
	if method == "GET" && path == "/" {
		conn.Write([]byte("HTTP/1.1 200 OK\r\nContent-Length: 12\r\n\r\nHello world!"))
	}  else if (method == "GET" && path == "/user-agent") {
		
	} else if (method == "GET" && path == "/validate-request") {
		conn.Write([]byte("HTTP/1.1 200 OK\r\nContent-Length: 12\r\n\r\nHello world!"))
	} else if (method == "GET" && path == "/echo") {
		dynamicPath := strings.Split(lines[0], " ")[1]
		conn.Write([]byte("HTTP/1.1 200 OK\r\nContent-Length: 12\r\n\r\n" + dynamicPath))
	} else {
		conn.Write([]byte("HTTP/1.1 404 Not Found\r\nContent-Length: 12\r\n\r\nNot found"))
	}

	conn.Close()
}
