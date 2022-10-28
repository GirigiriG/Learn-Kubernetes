package main

import (
	"fmt"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, req *http.Request) {
	SERVICE_DISCOVERY_PASSWORD := os.Getenv("SERVICE_DISCOVERY_PASSWORD")
	fmt.Fprintf(w, SERVICE_DISCOVERY_PASSWORD + "\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":3000", nil)
}
