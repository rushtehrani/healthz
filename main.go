package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
)

var (
	origin = flag.String("origin", "*", "Allowed origin.")
	port   = flag.Int("port", 5000, "Port number to serve /healthz (or customized path).")
)

func main() {
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
		w.Header().Set("Access-Control-Allow-Origin", *origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.WriteHeader(200)

		log.Println("/")
	})

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), nil))
}
