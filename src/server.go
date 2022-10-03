package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
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
		fmt.Println("A Client is connected")
		go handleClient(c, counter)
	}

}

func handleClient(c net.Conn, counter int) {

	//fmt.Println(c)
	//var client pk.Client
	//name := string(pk.EnterName())
	//client.Fiel = string(pk.EnterName())
	//client.Fiel = "archivo.txt"
	datatype1 := "salida" + strconv.Itoa(counter)

	//fmt.Println(client)

	b := make([]byte, 1000000)
	bs, err := c.Read(b)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		//datatype := string(b[:bs])
		fmt.Println("Bytes", bs)

		reader := bytes.NewReader(b)
		//dst := "/Programming/serverProyect/carpeta_salida/" + datatype
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

func main() {

	//src := "/Programming/serverProyect/archivos/yolov3.png"
	//dst := "/Programming/serverProyect/carpeta_salida/output.png"
	go servidor()

	//Copy(src, dst)
	var input string
	fmt.Scanln(&input)
}
