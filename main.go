package main

import (
	"fmt"
	"gmm"
	"net/http"
)

func main() {
	r := gmm.New()

	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "hello gmm")
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "%q : %q \n", k, v)
		}
	})

	r.Run(":3000")
}
