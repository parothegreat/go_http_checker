package main

import (
	//flag
	"fmt"
	"flag"
	"bufio"
	//other package
	"strings"
	"os"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	URL        string
	StatusCode int
	StatusText string
	Error      error
}

func checkURL(url string, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	
	resp, err := client.Get(url)

	if err != nil {
		results <- Result{URL: url, Error: err}
		return
	}

	defer resp.Body.Close()

	results <- Result{
		URL:        url,
		StatusCode: resp.StatusCode,
		StatusText: resp.Status,
		Error:      nil,
	}
}

func main() {
	fileName:= flag.String("file", "", "File path containing URLs (One URL per line")
	timeoutSeconds := flag.Int("timeout", 5, "Connection timeout in seconds.")

	flag.Parse()
	
	var urls []string

	if *fileName != ""{
		var err error
		urls, err = readURLsFromFile(*fileName)
        
        if err != nil {
            fmt.Printf("Failed to read file: %v\n", err)
            os.Exit(1)
		}
		fmt.Printf("Loaded %d URLs from file: %s\n", len(urls), *fileName)

	} else {
		fmt.Println("No file path specified. Using default hardcoded URLs.")
	    urls := []string{
		"https://google.com",
		"https://github.com/non-existent-page-404",
		"http://192.0.2.1:80",
	}
	 
 }  

	results := make(chan Result)
	var wg sync.WaitGroup

	fmt.Println("Starting concurrent checks...")

	for _, url := range urls {
		wg.Add(1)
		go checkURL(url, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	fmt.Println("--- Check Results ---")
	for res := range results {
		if res.Error != nil {
			fmt.Printf("%s: ERROR -> %v\n", res.URL, res.Error)
			continue
		}

		status := ""
		if res.StatusCode >= 200 && res.StatusCode < 300 {
			status = "OK"
		} else if res.StatusCode >= 400 && res.StatusCode < 500 {
			status = "Client Error"
		} else if res.StatusCode >= 500 {
			status = "Server Error"
		} else {
			status = "Other"
		}

		fmt.Printf("%s %s (%d %s)\n", status, res.URL, res.StatusCode, res.StatusText)
	}

	fmt.Println("---------------------")
	fmt.Println("Done. All results processed.")
}