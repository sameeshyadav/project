package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	log.Println("Server is running!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Inside / request handler and methord type is", r.Method)
		if r.Method == "GET" {
			http.Error(w, "ops bad request", http.StatusBadRequest)
			return
		}
		d, err := io.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "ops bad request", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "The output data for customer will be %s", d)
	})

	http.HandleFunc("/sameesh", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Inside /sameesh request handler and methord type is ", r.Method)
		if r.Method == "GET" {
			http.Error(w, "ops bad request", http.StatusBadRequest)
			return
		}
		d, err := io.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "ops bad request", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "The output data for customer will be %s from sameesh", d)
	})

	http.ListenAndServe(":9090", nil)
}
