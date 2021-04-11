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

	// Wait and repeated calls infinitive using checkWeb function
	for l := range c {
		// IMPORTANT: If this directive is included, yo wait per every call and this action "kill" the parallelism
		// time.Sleep(1 * time.Second)
		go checkWeb(l, c)
	}
}

func checkWeb(url string, c chan string) {
	time.Sleep(1 * time.Second)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(url + " - Error")
		c <- url
		return 
	}
	fmt.Println(url + " - " + strconv.Itoa(resp.StatusCode))
	c <- url
}