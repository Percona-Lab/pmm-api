package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addrF := flag.String("addr", "127.0.0.1:8080", "Address to listen")
	flag.Parse()

	log.Printf("Starting server on http://%s/ ...", *addrF)

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(*addrF, nil)
}
