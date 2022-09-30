package client

import "fmt"

type Client struct {
	Name       string
	Credential int
}

func EnterName() string {
	var nameFile string
	fmt.Println("Enter the name of the file, for example (hello.txt o image.jpg): ")
	fmt.Scan(&nameFile)
	return nameFile

}
