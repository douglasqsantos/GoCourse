package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// for i := 0; i < len(links); i++ {
	// fmt.Println(<-c)
	// }
	// for {
	// 	go checkLink(<-c, c)
	// }

	for l := range c {
		// go checkLink(l, c)
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("[!!] %v, might be down\n", link)
		c <- link
		return
	}

	fmt.Printf("[+] %v, is up!\n", link)
	c <- link
}
