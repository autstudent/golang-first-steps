package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func main()  {
	
	webs := []string {
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, url := range webs {
		go checkWeb(url, c)
	}

	for l := range c {
		// This func it is called with the same l argument everytime
		go func() {
			time.Sleep(1 * time.Second)
			checkWeb(l, c)
		}()
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