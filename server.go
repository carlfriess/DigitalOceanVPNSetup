package main

import (
	"io"
	"net/http"
	"log"
	"os"
	"time"
)

func exit() {
	time.Sleep(time.Second)
	os.Exit(0)
}

func handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("content-type", "application/json")
	io.WriteString(w, "{\"success\":true}")
	go exit()
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}