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

	handler := &QuestionHandler{
		db: dal,
	}

	mux := powermux.NewServeMux()

	mux.Route("/").OptionsFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Method", "GET,POST")
		w.Header().Add("Access-Control-Allow-Header", "Authorization")
	}).MiddlewareFunc(func(w http.ResponseWriter, r *http.Request, n func(w http.ResponseWriter, _ *http.Request)) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		n(w, r)
	})

	mux.Route("/").MiddlewareFunc(func(w http.ResponseWriter, r *http.Request, n func(w http.ResponseWriter, _ *http.Request)) {
		authToken := r.Header.Get("Authorization")
		if authToken == "sfu" {
			n(w, r)
			return
		}

		w.WriteHeader(http.StatusUnauthorized)
	})

	mux.Route("/ping").GetFunc(func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "pong")
	})

	questionRoute := mux.Route("/questions")

	questionRoute.GetFunc(handler.GetAllQuestions)
	questionRoute.PostFunc(handler.AddQuestion)
	questionRoute.Route("/:id").
		GetFunc(handler.GetQuestion).
		Route("/vote").
		PostFunc(handler.Upvote)

	port := os.Getenv("PORT")

	http.ListenAndServe(":"+port, mux)
}
