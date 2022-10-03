package myPackage

import (
	"encoding/gob"
	"fmt"
	"net"
)

type Client struct {
	Channel string
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

/*Solicitar datos de usuario*/
/*read := bufio.NewReader(os.Stdin)
fmt.Println("Input the channel: ")
channel, _ := read.ReadString('\n')

persona := InfoPersona{
	Nombre: "channel: " + channel
}
go Clients(persona)
var input string
fmt.Scanln(&input)
*/
