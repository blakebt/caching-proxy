package main

import (
	"io"
	"log"
	"net/http"
)

func getResponseBody(resp *http.Response) string {
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(bodyBytes)
}
