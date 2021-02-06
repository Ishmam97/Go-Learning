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
	// st := time.Now()
	// for _, link := range links {
	// 	seqCheckLink(link)
	// }
	// et := time.Now()
	// fmt.Println("time taken for sequential = ", et.Sub(st))
	//go routine
	st := time.Now()
	for _, link := range links {
		go seqCheckLink(link)
	}
	et := time.Now()
	fmt.Println("time taken for go Routine = ", et.Sub(st))
}

//sequential

func seqCheckLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + "might be down!")
		return
	}
	fmt.Println(link, " is up")
}
