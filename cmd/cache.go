package main

import "net/http"

type Response struct {
	Header http.Header
	Body   string
}

type ResponseCache map[string]Response
