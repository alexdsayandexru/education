package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := get("https://localhost:4433/random")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	reader := bufio.NewReader(resp.Body)

	events := make(chan string)

	go loop(events, reader)

	for event := range events {
		fmt.Print(event)
	}
}

func loop(events chan string, reader *bufio.Reader) {
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			events <- err.Error()
			close(events)
			break
		} else if len(line) > 1 {
			events <- string(line)
		}
	}
}

func get(rawurl string) (*http.Response, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // <--- Problem
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get(rawurl)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("got response status code %d\n", resp.StatusCode)
	}

	return resp, nil
}
