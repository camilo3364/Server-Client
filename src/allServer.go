package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"strings"

	pk "server.com/serverClient/src/myPackage"
)

func server() {
	var counter int
	s, err := net.Listen("tcp", ":9999")

	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := s.Accept()
		if err != nil {
			fmt.Println(err)

			continue
		}
		counter += 1
		go Cliente(c)
		fmt.Println("A client is connected")

	}

}
func Cliente(c net.Conn) {
	var person pk.Client
	err := gob.NewDecoder(c).Decode(&person)
	if err != nil {
		fmt.Println(err)
		return
	} else {

		if strings.Contains(person.Channel, "1") == true {
			fmt.Println("You are in the channel 1")
		} else if strings.Contains(person.Channel, "2") == true {
			fmt.Println("You are in the channel 2")
		} else {
			fmt.Println("You aren't in any channel")
		}

	}

}

func main() {

	go server()
	var input string
	fmt.Scanln(&input)
}
