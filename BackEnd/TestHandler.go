package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside / request handler")
	d, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "ops bad request", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "The output data for customer will be %s", d)
}
