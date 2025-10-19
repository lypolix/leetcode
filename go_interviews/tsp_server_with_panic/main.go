package main

import (
	"errors"
	"log"
	"net"

)

func main()  {
	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}

		go ClientHandler(conn)
	}

}
	

func ClientHandler(c net.Conn) {
	defer func ()  {
		if v := recover(); v != nil {
			log.Println("captured panic:", v)
		}
		c.Close()
	}()

	panic(errors.New("internal error"))
}
