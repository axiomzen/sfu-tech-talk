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

	mux.Route("/").MiddlewareFunc(func(w http.ResponseWriter, r *http.Request, n func(w http.ResponseWriter, r *http.Request)) {
		w.Header().Add("Access-Control-Allow-Origin", "https://sfu-talk.axiomzen.co")

		if r.Method == http.MethodOptions {
			w.Header().Add("Access-Control-Allow-Methods", "GET,POST")
			w.Header().Add("Access-Control-Allow-Headers", "Authorization")
			return
		}

		n(w, r)
	})

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

	questionRoute := mux.Route("/questions")
	questionRoute.GetFunc(handler.GetQuestions)
	questionRoute.PostFunc(handler.AddQuestion)
	questionRoute.Route("/:id/vote").PostFunc(handler.Upvote)
	// /questions/:id/vote

	port := os.Getenv("PORT")

	http.ListenAndServe(":"+port, mux)
}
