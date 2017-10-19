package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type QuestionHandler struct {
	db *DAL
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

}

func (h *QuestionHandler) Upvote(w http.ResponseWriter, r *http.Request) {

}
