package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
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
		counter += 1
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

	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	ch1 <- "5555"
	ch2 <- "8080"

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
			c, err := net.Dial("tcp", ":8080")
			if err != nil {
				fmt.Println(err)
				return
			}
			c.Write([]byte(split1[1]))
			c.Close()
			//
			fmt.Println("Ingresamos al chanel 1")
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
			c.Close()
		}

	}
}

func main() {

	go servidor()

	//Copy(src, dst)
	var input string
	fmt.Scanln(&input)
}
