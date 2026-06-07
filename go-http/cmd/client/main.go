package main

import (
	"fmt"
	"go-http/internal/shared"
	"io"
	"log"
	"net/http"
)

func apiCaller() {
	url := "https://jsonplaceholder.typicode.com/todos"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("status code", resp.StatusCode)
	fmt.Println("status", resp.Status)
}

func apiCallerAndReadBody() {
	url := "https://jsonplaceholder.typicode.com/todos"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.Status)
		return
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	bodyText := string(bodyBytes)

	max := 250
	if len(bodyText) < max {
		max = len(bodyText)
	}

	fmt.Println(bodyText[:max])
}

func main() {
	apiCaller()
	apiCallerAndReadBody()
	resp, err := shared.ApiCallerWithUnmarshal()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Response: Fact: %s, Length: %d\n", resp.Fact, resp.Length)
}
