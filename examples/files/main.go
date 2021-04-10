package main

import (
	"fmt"
	"os"
	"io"
)

func main()  {

	// Receive file name by args
	fname  := os.Args[1]

	// Read a file
	file, err := os.Open(fname)
	if err != nil {
		fmt.Println("Error Opening file:", err)
		os.Exit(1)
	}

	// Create an empty []byte to allocate file data
	bs := make([]byte, 99999)
	file.Read(bs)
	// Print file content
	fmt.Println(string(bs))

	// Print file content Stdout
	io.Copy(os.Stdout, file)

	// Close opened file
	if err := file.Close(); err != nil {
		fmt.Println("Error Closing file:", err)
		os.Exit(1)
	}
	
}
