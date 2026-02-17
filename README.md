
# Caching Proxy

## Overview
A simple caching proxy server written in Go that takes requests to locahost, forwards them to your origin of choice, and caches the responses
allowing for quick access when needed

### How it works

- Starts a server on the port of your choice ( Default: :4000 )
- Accepts origin of your choice ( Default: dummyjson.com )
- Forwards requests made to localhost to the provided origin
- Caches the response headers and response bodies for fast access
- Prints the response headers and response bodies to the terminal in a safe way to prevent terminal escape sequence issues

## Run Instructions

- Clone the repo or download the .Zip
- If on Windows: run caching.exe to start the server
- If on Linux: run ./caching to start the server
- Make requests to localhost:port
- Ex. curl http:localhost:4000/products forwards to https://dummyjson.com/products and returns a json object of products
