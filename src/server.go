package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

func servidor() {
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
		go handleClient(c)
	}
}

func handleClient(c net.Conn) {

	//var client
	//client.Fiel = client.EnterName()
	//client.Fiel = "archivo.txt"
	//fmt.Println(client)

	b := make([]byte, 1000000)
	bs, err := c.Read(b)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		datatype := string(b[:bs])
		fmt.Println("Bytes", bs, datatype)

		reader := bytes.NewReader(b)
		//dst := "/Programming/serverProyect/carpeta_salida/" + datatype
		out, err := os.Create("/Programming/serverProyect/carpeta_salida/" + datatype)
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
