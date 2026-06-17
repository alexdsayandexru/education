package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const PORT = 8888

func main() {
	http.Handle("/", http.FileServer(http.Dir("client")))
	http.HandleFunc("/random", random)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil))
}

func random(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Content-Type", "text/event-stream")

	for {
		value := rand.Intn(100)
		if _, err := fmt.Fprintf(w, "data: %s:%d \n\n", getHostname(), value); err == nil {
			fmt.Println(getHostname(), value)
		} else {
			fmt.Println(getHostname(), err)
			break
		}
		w.(http.Flusher).Flush()
		time.Sleep(2 * time.Second)
	}
}

func getHostname() string {
	name, _ := os.Hostname()
	return "[" + name + "] "
}
