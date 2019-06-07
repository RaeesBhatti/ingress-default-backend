package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Starting default backend on :"+os.Getenv("PORT"))
	http.HandleFunc("/healthz", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("OK!"))
	})
	log.Fatalln(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
