package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/andrewburian/powermux"
)

type QuestionHandler struct {
	dal *DAL
}

func (h *QuestionHandler) Setup(r *powermux.Route) {
	r.Route("/question").PostFunc(h.AddQuestion)
	r.Route("/vote/:id").PostFunc(h.Upvote)
	r.Route("/all").GetFunc(h.GetAllQuestions)
	r.Route("/question/:id").GetFunc(h.GetQuestion)
}

func decode(obj interface{}, body io.Reader) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(obj)
	return err
}

func encode(obj interface{}, writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(obj)
	return err
}

func (h *QuestionHandler) AddQuestion(w http.ResponseWriter, r *http.Request) {
	q := &Question{}

	if err := decode(q, r.Body); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if err := h.dal.AddQuestion(q); err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if err := encode(q, w); err != nil {
		http.Error(w, "Encoding error", http.StatusInternalServerError)
		return
	}
}

func (h *QuestionHandler) GetQuestion(w http.ResponseWriter, r *http.Request) {
	questionId := powermux.PathParam(r, "id")

	if questionId == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	questionInt, err := strconv.Atoi(questionId)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	q, err := h.dal.GetQuestion(questionInt)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	if err := encode(q, w); err != nil {
		http.Error(w, "Encoding error", http.StatusInternalServerError)
		return
	}

}

func (h *QuestionHandler) GetAllQuestions(w http.ResponseWriter, r *http.Request) {
	q, err := h.dal.GetAllQuestions()
	if err != nil {
		http.Error(w, "Database Error", http.StatusInternalServerError)
		return
	}

	if err := encode(q, w); err != nil {
		http.Error(w, "Encoding error", http.StatusInternalServerError)
		return
	}
}

func (h *QuestionHandler) Upvote(w http.ResponseWriter, r *http.Request) {
	questionId := powermux.PathParam(r, "id")

	if questionId == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	questionInt, err := strconv.Atoi(questionId)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err = h.dal.AddVote(questionInt)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
