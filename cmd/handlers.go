package main

import (
	"fmt"
	"net/http"
)

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

			fullResponse := Response{
				Header: resp.Header,
				Body:   getResponseBody(resp),
			}

			responseCache[forwardReq] = fullResponse
			fmt.Printf("%s\n\n", fullResponse.Header)
			fmt.Printf("%s\n", fullResponse.Body)

		} else {
			if cachedResp.Header.Get("X-Cache") != "HIT" {
				cachedResp.Header.Set("X-Cache", "HIT")
			}

			fmt.Printf("%s\n\n", cachedResp.Header)
			fmt.Printf("%s\n", cachedResp.Body)
		}

	}
}
