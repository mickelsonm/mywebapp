package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/go-martini/martini"
)

var (
	listenAddr = flag.String("http", ":8080", "http listen address")
)

func main() {
	flag.Parse()

	m := martini.Classic()

	m.Use(martini.Recovery())
	m.Use(martini.Logger())
	m.Use(martini.Static("public"))

	srv := &http.Server{
		Addr:         *listenAddr,
		Handler:      m,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("Starting server on 127.0.0.1%s\n", *listenAddr)
	log.Fatal(srv.ListenAndServe())
}
