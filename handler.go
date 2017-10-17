package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/andrewburian/powermux"
)

type QuestionHandler struct {
	db *DAL
}

func decode(r io.Reader, obj interface{}) error {
	decoder := json.NewDecoder(r)
	err := decoder.Decode(obj)
	return err
}

func render(w io.Writer, obj interface{}) error {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(obj)
	return err
}

func (h *QuestionHandler) GetQuestions(w http.ResponseWriter, r *http.Request) {

	questions, err := h.db.GetQuestions()
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if err := render(w, questions); err != nil {
		http.Error(w, "Render error", http.StatusInternalServerError)
		return
	}
}

func (h *QuestionHandler) AddQuestion(w http.ResponseWriter, r *http.Request) {

	question := &Question{}

	if err := decode(r.Body, question); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if err := h.db.AddQuestion(question); err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if err := render(w, question); err != nil {
		http.Error(w, "Render error", http.StatusInternalServerError)
		return
	}
}

func (h *QuestionHandler) Upvote(w http.ResponseWriter, r *http.Request) {

	id := powermux.PathParam(r, "id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if err := h.db.Upvote(idInt); err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
