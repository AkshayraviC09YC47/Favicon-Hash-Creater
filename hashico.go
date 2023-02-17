package main

import (
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("[!]Please provide a valid URL, If you are facing some issue please add http:// or https:// befor the url")
		fmt.Println("[+]Example: go run Hash.go https://www.google.com/favicon.ico")
		return
	}

	response, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Printf("[!]Error fetching URL: %v", err)
		return
	}
	defer response.Body.Close()

	favicon, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("[!]Error reading response body: %v", err)
		return
	}

	hash := fnv.New32a()
	hash.Write(favicon)

	fmt.Printf("[+]shodan search query: http.favicon.hash:%v\n", hash.Sum32())
}
