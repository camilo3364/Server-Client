package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

func client(src, dataType string, wg *sync.WaitGroup) {

	filerc, err := os.Open(src + dataType)
	//print(src + dataType)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer filerc.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(filerc)
	contents := buf.String()

	//fmt.Print(contents)
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
	defer wg.Done()
	c.Close()

}

func main() {
	//variables
	for {
		var wg sync.WaitGroup
		var name string
		src := "/Programming/codigos_go/serverClient/src/fields/"

		//Message
		fmt.Println("Please write the name of the field for exampe (image.png or hello.txt)")
		fmt.Scan(&name)

		var input string
		fmt.Scanln(&input)

		dataType := string(name)
		fmt.Println(dataType)

		wg.Add(1)
		go client(src, dataType, &wg)
		wg.Wait()
		fmt.Println("Do you need send other field? Yes or No")
		fmt.Scan(&name)

		if strings.ToLower(string(name)) == "yes" {
			print("Ingreso al if")
			main()
		} else if strings.ToLower(string(name)) == "no" {
			break
		}
		print(name)

	}

}
