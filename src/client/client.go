package main

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

/*
In this function there are two channels Ch1 and Ch2, which by means of a random number
between 1 and 9 assign the TCP port for data reception, in this way each time a client
is started it will have a different TCP port, because webSockets was not used.
*/
func listen(channel string) {
	var counter int

	rand.Seed(int64(time.Now().UnixNano()))
	randomNumber2 := rand.Intn(10)

	//fmt.Println("El number is: 555" + strconv.Itoa(randomNumber2))

	for {
		datatype1 := "salida" + strconv.Itoa(randomNumber2)
		b := make([]byte, 1000000)

		if channel == "1" {
			s, err := net.Listen("tcp", ":555"+strconv.Itoa(randomNumber2))

			if err != nil {

				main()
				continue
			}

			c, err := s.Accept()

			if err != nil {
				fmt.Println(err)
				continue
			}
			bs, err1 := c.Read(b)
			if err1 != nil {
				fmt.Println(err)
				return
			} else {
				fmt.Println("You receive a file by ch1")
				//fmt.Println("Bytes", bs)
				reader2 := string(b[:bs])
				split1 := strings.Split(reader2, "file...")
				reader := bytes.NewReader([]byte(split1[1]))

				//Decode the file type

				fmt.Println("The file that you receive is a: " + split1[0])
				out, err1 := os.Create("/Programming/codigos_go/serverClient/src/output/" + datatype1 + "." + split1[0])
				if err1 != nil {
					fmt.Println(err1)
				}
				defer out.Close()

				//Save the file in a different folder to ch2
				_, err1 = io.Copy(out, reader)
				if err1 != nil {
					fmt.Println(err1)
				}

				counter += 1

			}

			s.Close()
		} else if channel == "2" {
			//This is channel 2 and the TCP connection is different to ch1
			r, err := net.Listen("tcp", ":553"+strconv.Itoa(randomNumber2))
			if err != nil {
				main()
				continue

			}
			n, err := r.Accept()
			if err != nil {
				fmt.Println(err)
				continue
			}
			bs, err1 := n.Read(b)
			if err1 != nil {
				fmt.Println(err)
				return
			} else {
				reader2 := string(b[:bs])
				split1 := strings.Split(reader2, "file...")
				reader := bytes.NewReader([]byte(split1[1]))

				//Decode the file type

				fmt.Println("The file that you receive is a: " + split1[0])
				out, err1 := os.Create("/Programming/codigos_go/serverClient/src/output2/" + datatype1 + "." + split1[0])
				if err1 != nil {
					fmt.Println(err1)
				}
				defer out.Close()

				_, err1 = io.Copy(out, reader)
				if err1 != nil {
					fmt.Println(err1)
				}
				fmt.Println("You receive a file by ch2")
				counter += 1
			}
			r.Close()
		}

	}

}

/*
With this function the file is sent to the server when the client
is not in listening mode but in sending data mode.
*/

func client(src, dataType string, wg *sync.WaitGroup, channel string) {
	//Open the file
	filerc, err := os.Open(src + dataType)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer filerc.Close()

	//Convert the file in to bits
	buf := new(bytes.Buffer)
	buf.ReadFrom(filerc)
	contents := buf.String()

	//send file between client and server
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	//Send the file with extra bits with the type information and channel information.
	c.Write([]byte(channel))
	c.Write([]byte(contents))
	//Use a Waitgroup by go concurrency
	defer wg.Done()
	c.Close()

}

func main() {
	for {
		//variables
		var wg sync.WaitGroup
		var name string
		var channel string
		var channelOut string
		var listen1 string
		src := "/Programming/codigos_go/serverClient/src/fields/"

		fmt.Println("You need to shoose if receive(1) or send (2) a file: ")
		fmt.Scan(&listen1)
		//Do you need listen o write the message?
		fmt.Println("You need to join the channel (1 or 2): ")
		fmt.Scan(&channel)
		if channel == "1" {
			channelOut = "1chanel_..."
			fmt.Println(channel)
		} else if channel == "2" {
			channelOut = "2chanel_..."
		} else {
			fmt.Println("Please, you should select a channel ")
			main()
		}
		//Select the funtion if listen a message or send the message by channels
		if listen1 == "1" {
			for {
				fmt.Println("In this moment you are listen a message...")
				listen(channel)
			}

		} else if listen1 == "2" {

			//Message
			fmt.Println("Please write the name of the file for example (image.png or hello.txt): ")
			fmt.Scan(&name)

			var input string
			fmt.Scanln(&input)

			dataType := string(name)
			split1 := strings.Split(dataType, ".")
			fmt.Println("the file type is: " + split1[1])
			channelOut = channelOut + split1[1] + "file..."
			wg.Add(1)
			go client(src, dataType, &wg, channelOut)
			wg.Wait()
			fmt.Println("Do you need send other file? Yes or No")
			fmt.Scan(&name)

			//Know if you want to send a message again or if you want to exit.
			if strings.ToLower(string(name)) == "yes" {
				//print("Ingreso al if")
				main()
			} else if strings.ToLower(string(name)) == "no" {
				break
			} else {
				fmt.Println("Do you need send other file? Yes or No")
				continue
			}
		}

		//Channel suscribe

	}

}
