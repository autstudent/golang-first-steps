package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {}

func main()  {
	resp, err := http.Get("http://google.com")

	// Quit program when a new error appears
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Read http respon body
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// Log information in a terminal nicely
	// io.Copy(os.Stdout, resp.Body)
	
	// Log information with a custom writter
	lw := logWriter{}
	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte)  (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}