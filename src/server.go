package main

import (
	"fmt"

	pk "server.com/serverClient/src/client"
)

func main() {
	var client pk.Client
	client.Name = "Juan"
	client.Credential = 2
	fmt.Println(client)
}
