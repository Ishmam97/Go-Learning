package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type myWriter struct{}

func main() {
	res, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Printf("Error : %v", err)
		os.Exit(1)
	}
	io.Copy(myWriter{}, res.Body)
}
func (myWriter) Write(p []byte) (int, error) {
	fmt.Println(string(p), "\nLEN: ", len(p))
	return len(p), nil
}
