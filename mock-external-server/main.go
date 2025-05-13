package main

import (
	// "fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(32 * time.Second) // default timeouts 10 15 30
		// fmt.Fprintln(w, "allo")
		w.WriteHeader(500)

	})

	http.ListenAndServe(":8080", nil)
}
