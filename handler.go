package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type QuestionHandler struct {
	db *DAL
}

func encode(obj interface{}, w io.Writer) error {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(obj)
	return err
}

func decode(obj interface{}, r io.Reader) error {
	decoder := json.NewDecoder(r)
	err := decoder.Decode(obj)
	return err
}

func (h *QuestionHandler) GetQuestions(w http.ResponseWriter, r *http.Request) {

	allQuestions, err := h.db.GetQuestions()
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if err := encode(allQuestions, w); err != nil {
		http.Error(w, "Render error", http.StatusInternalServerError)
		return
	}

}

func (h *QuestionHandler) AddQuestion(w http.ResponseWriter, r *http.Request) {

	question := &Question{}
	if err := decode(question, r.Body); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if err := h.db.AddQuestion(question); err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if err := encode(question, w); err != nil {
		http.Error(w, "Render error", http.StatusInternalServerError)
		return
	}
}
