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
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    return nil
}
