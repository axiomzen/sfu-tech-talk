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

	questionHandler := &QuestionHandler{
		db: dal,
	}

	mux := powermux.NewServeMux()

	mux.Route("/").OptionsFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Method", "GET,POST")
		w.Header().Add("Access-Control-Allow-Header", "Authorization")
	}).MiddlewareFunc(func(w http.ResponseWriter, r *http.Request, n func(w http.ResponseWriter, _ *http.Request)) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		n(w, r)
	})

	mux.Route("/").MiddlewareFunc(func(w http.ResponseWriter, r *http.Request, n func(w http.ResponseWriter, r *http.Request)) {
		authToken := r.Header.Get("Authorization")
		if authToken == "sfu" {
			n(w, r)
			return
		}

		w.WriteHeader(http.StatusUnauthorized)
	})

	mux.Route("/ping").GetFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "pong")
	})

	mux.Route("/questions").
		GetFunc(questionHandler.GetQuestions).
		PostFunc(questionHandler.AddQuestion).
		Route("/:id").
		GetFunc(questionHandler.GetQuestion).
		Route("/vote").
		PostFunc(questionHandler.Upvote)

	port := os.Getenv("PORT")

	http.ListenAndServe(":"+port, mux)

}
