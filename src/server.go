package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"

	pk "server.com/serverClient/src/myPackage"
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
		fmt.Println("Las conexiones actuales son: " + strconv.Itoa(counter))
		fmt.Println("A Client sent a file")
		go handleClient2(c, counter)

	}

}

func handleClient(c net.Conn, counter int) {

	datatype1 := "salida" + strconv.Itoa(counter)

	b := make([]byte, 1000000)
	bs, err := c.Read(b)
	if err != nil {
		fmt.Println(err)
		return
	} else {

		fmt.Println("Bytes", bs)

		reader := bytes.NewReader(b)
		out, err := os.Create("/Programming/codigos_go/serverClient/src/output/" + datatype1)
		if err != nil {
			fmt.Println(err)
		}
		defer out.Close()

		_, err = io.Copy(out, reader)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func handleClient2(c net.Conn, counter int) {

	//datatype1 := "salida" + strconv.Itoa(counter)

	b := make([]byte, 1000000)
	bs, err := c.Read(b)
	if err != nil {
		fmt.Println(err)
		return
	} else {

		fmt.Println("Bytes", bs)

		//reader := bytes.NewReader(b)
		reader2 := string(b[:bs])
		fmt.Println(reader2)
		split1 := strings.Split(reader2, "chanel_...")
		split2 := strings.Split(reader2, "chanel_...")

		if split1[0] == ("1") {
			//send to message
			//send file between client and server
			for i := 0; i < 10; i++ {
				//fmt.Println("Ingresamos al for")
				c, err := net.Dial("tcp", ":555"+strconv.Itoa(i))
				if err != nil {
					continue
				}
				c.Write([]byte(split1[1]))
				c.Close()
				//
				//fmt.Println("Ingresamos al chanel 1")

			}
			//comienzo a comentar

			/*
				c, err := net.Dial("tcp", ":8080")
				if err != nil {
					fmt.Println(err)
					return
				}
				c.Write([]byte(split1[1]))
				defer c.Close()
				//
				fmt.Println("Ingresamos al chanel 1")

				//Other port for other connection
				//

				c1, err1 := net.Dial("tcp", ":8081")
				if err1 != nil {
					fmt.Println(err)
					return
				}
				c1.Write([]byte(split1[1]))
				defer c1.Close()
				//
				fmt.Println("Ingresamos al chanel 1.1")
			*/
			//termino de comentar

			//
			/*out, err := os.Create("/Programming/codigos_go/serverClient/src/output/" + datatype1)
			if err != nil {
				fmt.Println(err)
			}
			defer out.Close()

			_, err = io.Copy(out, reader)
			if err != nil {
				fmt.Println(err)
			}
			*/

		} else if split2[0] == ("2") {
			fmt.Println("Ingresamos al chanel2")
			c, err := net.Dial("tcp", ":5555")
			if err != nil {
				fmt.Println(err)
				return
			}
			c.Write([]byte(split2[1]))
			defer c.Close()
		}

	}
}

func main() {

	fmt.Println(pk.CreateToClient())

	go servidor()

	//Copy(src, dst)
	var input string
	fmt.Scanln(&input)
}
