package main

import (
	"flag"
	"fmt"
	"net/http"
)

type Application struct {
	responseCache ResponseCache
}

func main() {

	mux := http.NewServeMux()
	port := flag.String("port", ":4000", "HTTP network address")
	origin := flag.String(
		"origin",
		"http://dummyjson.com",
		"URL of server to forward requests",
	)

	responseCache := make(map[string]*http.Response)

	app := Application{
		responseCache: responseCache,
	}

	flag.Parse()
	mux.HandleFunc("/", requestHandler(*origin, app.responseCache))

	fmt.Printf("Starting server port=%s\n", *port)
	http.ListenAndServe(*port, mux)
}
