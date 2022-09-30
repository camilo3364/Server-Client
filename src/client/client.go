package client

import (
	"bytes"
	"fmt"
	"net"
	"os"
)

type Client struct {
	Name   string
	Serial int
}

func EnterName() string {
	var nameFile string
	fmt.Println("Enter the name of the file, for example (hello.txt o image.jpg): ")
	fmt.Scan(&nameFile)
	return nameFile

}

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
	EnterName()
	//c.Write([]byte(contents))
	//save the imgByte to file
	//out, err := os.Create("./QRImg.png")

	c.Close()

}

func main() {

	src := "/Programming/serverProyect/archivos/"
	dataType := "prueba1.txt"
	fmt.Println(dataType)
	//enterName()
	//go EnterName()
	go client(src, dataType)

	//f := "/Programming/serverProyect/archivos/prueba1.txt"
	//openFile(f)
	var input string
	fmt.Scanln(&input)

}
