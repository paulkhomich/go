package main

import (
	"net/http"
	"io"
	"io/ioutil"
	"os"
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Printf(<-ch)
	}
}

func fetch(url string, ch chan<-string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Error while fetching %s: %s\n", url, err)
		
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	ch <- fmt.Sprintf("%-8.2f\t%-8d\t%s\n", time.Since(start).Seconds(), nbytes, url)
}