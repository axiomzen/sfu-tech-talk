package main

import "github.com/go-pg/pg"

// Question model
type Question struct {
	ID      int    `json:"id"`
	Text    string `json:"body"`
	Author  string `json:"author"`
	Upvotes int    `json:"upvotes"`
}

// DAL struct
type DAL struct {
	conn *pg.DB
}

func (dal *DAL) GetQuestions() ([]*Question, error) {

	questions := make([]*Question, 0)

	err := dal.conn.Model(&questions).Select()

	return questions, err
}

func (dal *DAL) AddQuestion(q *Question) error {

	_, err := dal.conn.Model(q).Insert()

	return err
}

func (dal *DAL) Upvote(id int) error {

	_, err := dal.conn.Model(&Question{}).Where("id = ?", id).Set("upvotes = upvotes + 1").Update()

	return err
}
