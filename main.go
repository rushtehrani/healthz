package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
)

var (
	path   = flag.String("path", "/", "Custom path")
	origin = flag.String("origin", "*", "Allowed origin.")
	port   = flag.Int("port", 5000, "Port number to serve /healthz (or customized path).")
)

func main() {
	flag.Parse()

	http.HandleFunc(*path, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
		w.Header().Set("Access-Control-Allow-Origin", *origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.WriteHeader(http.StatusNoContent)

		log.Println("/")
	})

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), nil))
}
