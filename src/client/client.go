package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

func sendTcp(tcp string) {
	name := []byte(tcp)
	c, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(name)
	c.Write([]byte(tcp))

	c.Close()

}
func listen(channel string) {
	var counter int
	for {
		if channel == "1" {
			s, err := net.Listen("tcp", ":8080")
			if err != nil {

				main()
				continue

				//fmt.Println(err)
				//fmt.Println("Ocurrio error en 1")
				//return
			}
			c, err := s.Accept()
			if err != nil {
				//main()
				fmt.Println(err)
				fmt.Println("Ocurrio error en 2")
				continue
			}
			counter += 1
			fmt.Println("You receive a file by ch1")
			fmt.Println(c)
			//go handleClient2(c, counter)
			s.Close()
		} else if channel == "2" {
			fmt.Println("You receive a file by ch2")
			r, err := net.Listen("tcp", ":5555")
			if err != nil {
				//r.Close()
				main()
				continue
				//fmt.Println(err)
				//fmt.Println("Ocurrio error en 3")

				//return

			}
			n, err := r.Accept()
			if err != nil {
				fmt.Println(err)
				fmt.Println("Ocurrio error en 4")

				continue
			}
			counter += 1
			fmt.Println("You receive a file by ch2")
			fmt.Println(n)
			//go handleClient2(c, counter)
			r.Close()
		}

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
		RandomCrypto, _ := rand.Prime(rand.Reader, 12)
		randomNumber := RandomCrypto.String()
		fmt.Println(randomNumber)
		var wg sync.WaitGroup
		var name string
		var channel string
		var channelOut string
		var listen1 string
		src := "/Programming/codigos_go/serverClient/src/fields/"

		fmt.Println("You need to shoose if receive(1) or send (2) a file: ")
		fmt.Scan(&listen1)
		//Do you need listen o write the message?
		fmt.Println("You need to join the channel (1 or 2): ")
		fmt.Scan(&channel)
		if channel == "1" {
			channelOut = "1chanel_..." + randomNumber + "_..."
			sendTcp(randomNumber)
			fmt.Println(randomNumber)
			fmt.Println(channel)
		} else if channel == "2" {
			channelOut = "2chanel_..." + randomNumber + "_..."
		} else {
			fmt.Println("Please, you should select a channel ")
			main()
		}

		if listen1 == "1" {
			//flag := false
			//var loop string
			fmt.Println("In this momento you are listen a message...")
			go listen(channel)
			/*for flag == false {
				fmt.Println("In this momento you are listen a message...")
				go listen(channel)
				fmt.Println("Do you want to receive a new file (1) or go out (2)")
				fmt.Scan(&loop)
				if loop == "1" {
					flag = true
				} else {
					defer listen(channel)
				}

			}*/

		} else if listen1 == "2" {

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

		//Channel suscribe

	}

}
