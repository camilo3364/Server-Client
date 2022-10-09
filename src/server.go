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

func listenTcp() {
	//var counter int
	s, err := net.Listen("tcp", ":8080")

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
		//counter += 1
		fmt.Println("A Client sent a tcp")
		go decodificTcp(c)
		tcp := decodificTcp(c)
		go servidor(tcp)

	}
}

func servidor(tcp string) {
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
		go handleClient2(c, counter, tcp)

	}

}
func decodificTcp(c net.Conn) (r string) {

	b := make([]byte, 1000000)

	bs, err := c.Read(b)

	if err != nil {
		fmt.Println(err)
		return

	} else {

		//fmt.Println("Bytes", bs)
		reader2 := string(b[:bs])
		fmt.Println("El tcp enviado por parte del cliente es: " + reader2)
		return reader2
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

func handleClient2(c net.Conn, counter int, tcp string) {
	var ch1, ch2, ch3 string
	ch3 = ch2
	ch2 = ch1
	ch1 = tcp

	fmt.Println(tcp, ch1, ch2, ch3)

	fmt.Println("Nos encontramos en el handle:" + tcp)
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
		split3 := strings.Split(split1[1], "_...")
		split4 := strings.Split(split2[1], "_...")

		fmt.Println(split3[0], split4[0])

		if split1[0] == ("1") {
			//send to message
			//send file between client and server
			c, err := net.Dial("tcp", ":"+tcp)
			//s, err1 := net.Dial("tcp", ":"+split4[0])
			if err != nil {
				fmt.Println(err)
				return
			}

			c.Write([]byte(split1[1]))
			//s.Write([]byte(split1[1]))
			//s.Close()
			c.Close()
			//
			fmt.Println("Ingresamos al chanel 1")
			defer listenTcp()
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
			c, err := net.Dial("tcp", ":"+tcp)
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

	//go servidor()
	go listenTcp()

	//Copy(src, dst)
	var input string
	fmt.Scanln(&input)
}
