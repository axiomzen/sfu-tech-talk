package main

import (
	"io"
	"net/http"
	"os"

	"github.com/andrewburian/powermux"
)

func main() {

	mux := powermux.NewServeMux()

	mux.Route("/ping").GetFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "pong")
	})

	port := os.Getenv("PORT")

	http.ListenAndServe(":"+port, mux)
}
