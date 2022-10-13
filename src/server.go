package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	//pk "server.com/serverClient/src/myPackage"
)

func servidor() {
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
		if c != nil {
			counter += 1
		}
		fmt.Println("There are : " + strconv.Itoa(counter) + " connections")
		go handleClient2(c, counter)
	}
}

func handleClient2(c net.Conn, counter int) {

	b := make([]byte, 1000000)
	bs, err := c.Read(b)
	if err != nil {
		fmt.Println(err)
		return
	} else {

		//fmt.Println("Bytes", bs)
		reader2 := string(b[:bs])
		//fmt.Println(reader2)
		split1 := strings.Split(reader2, "chanel_...")
		split11 := strings.Split(split1[1], "file...")
		split2 := strings.Split(reader2, "chanel_...")
		fmt.Println("the file type is : " + split11[0])

		if split1[0] == ("1") {
			//send to message
			//send file between client and server
			for i := 0; i < 10; i++ {
				c, err := net.Dial("tcp", ":555"+strconv.Itoa(i))
				if err != nil {
					continue
				}
				c.Write([]byte(split1[1]))
				c.Close()

			}

		} else if split2[0] == ("2") {
			for i := 0; i < 10; i++ {
				c, err := net.Dial("tcp", ":553"+strconv.Itoa(i))
				if err != nil {
					continue
				}
				c.Write([]byte(split2[1]))
				c.Close()

			}

		}

	}
}

func main() {

	go servidor()
	var input string
	fmt.Scanln(&input)
}
