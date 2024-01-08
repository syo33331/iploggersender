package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter host: ")
host, _ := reader.ReadString('\n')
host = strings.TrimSpace(host)

fmt.Print("Enter filename: ")
filename, _ := reader.ReadString('\n')
filename = strings.TrimSpace(filename)

url := "https://" + host + "/" + filename

	fmt.Print("Enter number of requests: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	numberOfRequests, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Error: Invalid number of requests")
		return
	}

	for i := 0; i < numberOfRequests; i++ {
		err := sendRequest(url)
		if err != nil {
			fmt.Printf("Request %d failed: %v\n", i+1, err)
			continue
		}
		fmt.Printf("Request %d successful\n", i+1)
	}
}

func sendRequest(url string) error {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "ja,en-US;q=0.7,en;q=0.3")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("TE", "trailers")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/4.0 (Windows NT 10.0; Win64; x64; rv:121.0) Gecko/20100101 Firefox/100.0")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
