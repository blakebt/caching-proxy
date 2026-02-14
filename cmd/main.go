package main

import (
	"flag"
	"fmt"
	"io"
	"log"
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

func requestHandler(origin string, responseCache ResponseCache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqURI := r.URL.RequestURI()
		forwardReq := fmt.Sprintf("%s%s?limit=3", origin, reqURI)

		fmt.Printf("\n\nMaking Request to %s\n\n", forwardReq)
		cachedResp, exists := responseCache[forwardReq]
		if !exists {
			resp, err := http.Get(forwardReq)
			if err != nil {
				// TODO
				// Clean this up to be more precise in error reporting
				http.Error(w, "Failed to get response", http.StatusBadRequest)
			}

			resp.Header.Add("X-Cache", "MISS")
			responseCache[forwardReq] = resp

			respBody := getResponseBody(resp)
			fmt.Println(respBody)

		} else {
			if cachedResp.Header.Get("X-Cache") != "HIT" {
				cachedResp.Header.Set("X-Cache", "HIT")
			}

			log.Println(cachedResp)
		}

	}
}

func getResponseBody(resp *http.Response) string {
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(bodyBytes)
}
