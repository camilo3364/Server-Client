package myPackage

import (
	"encoding/gob"
	"fmt"
	"net"
)

type Client struct {
	Channel string
	Port    string
}

func Clients(persona Client) {
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = gob.NewEncoder(c).Encode(persona)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}
func CreateToClient() string {

	client2 := Client{"ch1", ":8082"}
	//client3 := Client{"ch1", 8083}
	//client4 := Client{"ch2", 8084}
	//client5 := Client{"ch2", 8085}

	return client2.Port
}
