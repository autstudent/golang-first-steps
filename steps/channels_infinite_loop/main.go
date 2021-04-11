package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main()  {
	
	webs := []string {
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://amazon.com/pepepeep",
		"http://amazon.asdasd/pepepeep",
	}

	c := make(chan string)

	for _, url := range webs {
		go checkWeb(url, c)
	}

	// Wait and repeated calls infinitive using checkWeb function
	for {
		go checkWeb(<-c, c)
	}
}

func checkWeb(url string, c chan string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(url + " - Error")
		c <- url
		return 
	}
	fmt.Println(url + " - " + strconv.Itoa(resp.StatusCode))
	c <- url
}