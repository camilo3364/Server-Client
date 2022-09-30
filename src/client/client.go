package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
)

func client(src, dataType string) {

	filerc, err := os.Open(src + dataType)
	print(src + dataType)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer filerc.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(filerc)
	contents := buf.String()

	fmt.Print(contents)
	//send file between client and server
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	c.Write([]byte(contents))
	//
	//c.Write([]byte(contents))
	//save the imgByte to file
	//out, err := os.Create("./QRImg.png")

	c.Close()

}

func main() {
	//var name string
	//fmt.Scan(&name)
	src := "/Programming/serverProyect/archivos/"
	dataType := "prueba1.txt"
	fmt.Println(dataType)
	go client(src, dataType)

	var input string
	fmt.Scanln(&input)

}
