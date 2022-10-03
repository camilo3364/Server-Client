package myPackage

import "fmt"

type Client struct {
	Fiel string
}

func EnterName() string {
	var nameFile string
	fmt.Println("Enter the name of the file, for example (hello.txt o image.jpg): ")
	fmt.Scan(&nameFile)
	return nameFile

}
