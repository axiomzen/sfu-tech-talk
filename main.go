package main

import (
	"io"
	"net/http"
	"os"

	"github.com/go-pg/pg"

	"github.com/andrewburian/powermux"
)

func main() {

	dbOpts, err := pg.ParseURL(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	dal := &DAL{
		conn: pg.Connect(dbOpts),
	}

	mux := powermux.NewServeMux()

	mux.Route("/ping").GetFunc(func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "pong")
	})

	port := os.Getenv("PORT")

	http.ListenAndServe(":"+port, mux)
}
