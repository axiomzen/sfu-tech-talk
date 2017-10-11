package main

import (
	"github.com/go-pg/pg"
)

// Question that is asked
type Question struct {
	ID      int    `json:"id"`
	Text    string `json:"text"`
	Upvotes int    `json:"upvotes"`
	Author  string `json:"author"`
}

type DAL struct {
	conn *pg.DB
}

func (db *DAL) AddQuestion(q *Question) error {
	_, err := db.conn.Model(q).Insert()

	return err
}

func (db *DAL) AddVote(id int) error {
	_, err := db.conn.Model(&Question{}).Set("upvotes = upvotes + 1").Where("id = ?", id).Update()

	return err
}

func (db *DAL) GetQuestion(id int) (*Question, error) {
	question := &Question{}
	if err := db.conn.Model(question).Where("id = ?", id).Select(); err != nil {
		return nil, err
	}
	return question, nil
}

func (db *DAL) GetAllQuestions() ([]*Question, error) {
	questionsArr := make([]*Question, 0)
	if err := db.conn.Model(questionsArr).Select(); err != nil {
		return nil, err
	}
	return questionsArr, nil
}
