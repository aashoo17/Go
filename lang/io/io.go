package io

import (
	"fmt"
	"os"
)

//write to a file

func WriteToFile() {
	//a string which we want to write
	sl := "Hello World"
	//open a file for writing, create if file not available
	f, _ := os.Create("file.txt")
	defer f.Close()

	//since write() takes []byte slice use type conversion to turn string to []byte
	f.Write([]byte(sl))
}

//read from a file
func ReadFromFile() {
	//create a slice of size 100 byte
	sl := make([]byte, 100)
	//open file for reading only
	f, _ := os.Open("file.txt")
	defer f.Close()
	//read byte in sl
	f.Read(sl)
	//convert []byte to string and print
	fmt.Println(string(sl))
}

//TODO: open file using OpenFile() for any flags like reading,writing,appending and all possible

func OpenFile() {
	f, _ := os.OpenFile("file.txt", os.O_RDWR|os.O_CREATE, 0755)
	defer f.Close()

	f.Write([]byte("this is what we want to write"))
}
