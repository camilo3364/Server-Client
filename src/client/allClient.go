package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"

	pk "server.com/serverClient/src/myPackage"
)

/*func listen() {
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
		go handleClient(c, counter)
	}

}
*/

func conect(person pk.Client) {
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = gob.NewEncoder(c).Encode(person)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}

func sent(src, dataType string, wg *sync.WaitGroup) {
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

/*func handleClient(c net.Conn, counter int) {

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
*/

func main() {

	for {

		//variables
		var wg sync.WaitGroup
		var name string
		src := "/Programming/codigos_go/serverClient/src/fields/"

		//Input the channel of client
		var person pk.Client
		fmt.Println("Please write the channel that you need join (1,2):")
		fmt.Scan(&name)
		person.Channel = name
		go conect(person)

		//Message
		fmt.Println("Please write the name of the file for exampe (image.png or hello.txt): ")
		fmt.Scan(&name)

		var input string
		fmt.Scanln(&input)

		dataType := string(name)

		wg.Add(1)
		go sent(src, dataType, &wg)
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
