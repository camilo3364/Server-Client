package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

/*
	func connectChannel(channel string) {
		channel1 := "111111111."
		//channel2 := "222222222."
		c, err := net.Dial("tcp", ":9999")
		if err != nil {
			fmt.Println(err)
			return
		}
		c.Write([]byte(channel1))
		c.Write([]byte(channel))

		c.Close()
	}
*/

func listen() {
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
		fmt.Println("A Client sent a file")
		fmt.Println(c)
		//go handleClient2(c, counter)
	}

}
func client(src, dataType string, wg *sync.WaitGroup, channel string) {

	filerc, err := os.Open(src + dataType)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer filerc.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(filerc)
	contents := buf.String()

	//send file between client and server
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	c.Write([]byte(channel))
	c.Write([]byte(contents))

	defer wg.Done()
	c.Close()

}

func main() {

	for {
		//variables
		var wg sync.WaitGroup
		var name string
		var channel string
		var channelOut string
		src := "/Programming/codigos_go/serverClient/src/fields/"

		//Channel suscribe
		fmt.Println("You need to join the channel (1 or 2): ")
		fmt.Scan(&channel)
		if channel == "1" {
			channelOut = "1chanel_..."
		} else if channel == "2" {
			channelOut = "2chanel_..."
		} else {
			fmt.Println("Please, you should select a channel ")
			main()
		}

		//go connectChannel(channel)
		//Message
		fmt.Println("Please write the name of the file for example (image.png or hello.txt): ")
		fmt.Scan(&name)

		var input string
		fmt.Scanln(&input)

		dataType := string(name)

		wg.Add(1)
		go client(src, dataType, &wg, channelOut)
		wg.Wait()
		fmt.Println("Do you need send other file? Yes or No")
		fmt.Scan(&name)

		if strings.ToLower(string(name)) == "yes" {
			print("Ingreso al if")
			main()
		} else if strings.ToLower(string(name)) == "no" {
			break
		} else {
			fmt.Println("Do you need send other file? Yes or No")
			continue
		}

	}

}
