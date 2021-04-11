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

	for i := 0; i < len(webs); i++ {
		fmt.Println(<-c)
	}
}

func checkWeb(url string, c chan string) {
	resp, err := http.Get(url)
	if err != nil {
		c <- url + " - Error"
		return 
	}
	c <- url + " - " + strconv.Itoa(resp.StatusCode)
}