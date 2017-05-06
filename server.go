package main

import (
	"io"
	"net/http"
	"log"
	"os"
	"time"
	"github.com/rs/cors"
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

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	// Support CORS:
	corsHandler := cors.Default().Handler(mux)

	log.Fatal(http.ListenAndServe(":80", corsHandler))

}