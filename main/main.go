package main

import(
       "fmt"
	   "net/http"
	   "os"
	   "strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Penggunaan: go run main.go <URL>")
		fmt.Println("   Contoh: go run main.go google.com")
		os.Exit(1)
	}

	!strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url 
	}
fmt.Printf("Checking: %s\n", url)
resp, err := http.Get(url)

if err != nil {
		fmt.Printf("Failed to request (URL not valid or network issues): %s\n", err)
		return
	}
	defer resp.Body.Close()

	statusCode := resp.StatusCode
	statusText := resp.Status

	
}