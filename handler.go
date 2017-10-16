package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type QuestionHandler struct {
	db *DAL
}

func decode(obj interface{}, r io.Reader) error {
	decoder := json.NewDecoder(r)
	err := decoder.Decode(obj)
	return err
}

func encode(obj interface{}, w io.Writer) error {
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

	if err := encode(questions, w); err != nil {
		http.Error(w, "Rendering error", http.StatusInternalServerError)
		return
	}

}
