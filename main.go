package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("Error starting TCP server: %v", err)
	}
	defer listener.Close()
	log.Println("Server started on port 5001")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		// Handle the connection in a new goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	file, err := os.Open("Orbits_Mock_Session.txt")
	if err != nil {
		log.Printf("Failed to open file: %v", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	log.Printf("Handling new connection from %v", conn.RemoteAddr())

	for {
		time.Sleep(time.Second / 2)
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			log.Println("Reached end of file")
			break
		} else if err != nil {
			log.Printf("Error reading from file: %v", err)
			continue
		}

		if len(line) > 0 {
			log.Printf("Sending line: %s", line)
			_, err := conn.Write([]byte(line))
			if err != nil {
				log.Printf("Error writing to connection: %v", err)
				return
			}
		}
	}
	log.Printf("Connection closed from %v", conn.RemoteAddr())
}
