package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://ishmamsolaiman.com",
	}

	c := make(chan string)

	// st := time.Now()
	// for _, link := range links {
	// 	seqCheckLink(link)
	// }
	// et := time.Now()
	// fmt.Println("time taken for sequential = ", et.Sub(st))
	//go routine
	st := time.Now()
	for _, link := range links {
		go seqCheckLink(link, c)

	}
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
	et := time.Now()
	fmt.Println("time taken for go Routine = ", et.Sub(st))
}

//sequential

func seqCheckLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "down"
		return
	}
	fmt.Println(link, " is up")
	c <- "up"
}
