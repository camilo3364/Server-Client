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
	c.Write([]byte(contents))

	defer wg.Done()
	c.Close()

}

func main() {

	for {
		//variables
		var wg sync.WaitGroup
		var name string
		src := "/Programming/codigos_go/serverClient/src/fields/"

		//Message
		fmt.Println("Please write the name of the file for exampe (image.png or hello.txt): ")
		fmt.Scan(&name)

		var input string
		fmt.Scanln(&input)

		dataType := string(name)

		wg.Add(1)
		go client(src, dataType, &wg)
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

}
