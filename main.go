package main

import (
	"io"
	"net/http"
	"os"

	"github.com/andrewburian/powermux"
)

func main() {

	mux := powermux.NewServeMux()

	mux.Route("/").MiddlewareFunc(func(w http.ResponseWriter, r *http.Request, n func(w http.ResponseWriter, r *http.Request)) {
		authToken := r.Header.Get("Authorization")
		if authToken != "sfu" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		n(w, r)
	})

	mux.Route("/ping").GetFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "pong")
	})

	port := os.Getenv("PORT")

	http.ListenAndServe(":"+port, mux)

}
