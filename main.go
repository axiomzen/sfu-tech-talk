package main

import (
	"io"
	"net/http"
	"os"

	"github.com/andrewburian/powermux"

	"github.com/go-pg/pg"
)

func main() {

	dal := &DAL{}

	dbOpts, err := pg.ParseURL(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	dal.conn = pg.Connect(dbOpts)

	handler := &QuestionHandler{}
	handler.dal = dal

	mux := powermux.NewServeMux()

	mux.Route("/ping").GetFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "pong")
	})

	handler.Setup(mux.Route("/"))

	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, mux)

}
