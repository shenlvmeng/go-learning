package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

func main() {
	flag.Parse()
	cancel := make(chan struct{})
	responses := make(chan *http.Response)
	var wg sync.WaitGroup

	for _, url := range flag.Args() {
		go func(url string) {
			wg.Add(1)
			defer wg.Done()
			req, err := http.NewRequest("HEAD", url, nil)
			if err != nil {
				fmt.Printf("HEAD %s: %v\n", url, err)
				return
			}
			req.Cancel = cancel
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				fmt.Print("request %s: %v\n", url, err)
				return
			}
			responses <- res
		}(url)
	}

	res := <-responses
	defer res.Body.Close()
	close(cancel)
	fmt.Println("done", res.Request.URL)
	for key, val := range res.Header {
		fmt.Printf("%s: %s, ", key, strings.Join(val, ", "))
	}
	wg.Wait()
}
