package main

import "github.com/go-pg/pg"

// Question model
type Question struct {
	id      int    `json:"id"`
	text    string `json:"text"`
	upvotes int    `json:"upvotes"`
	author  string `json:"author"`
}

type DAL struct {
	conn *pg.DB
}

func (dal *DAL) GetQuestion(id int) (*Question, error) {
	question := &Question{}
	err := dal.conn.Model(question).Where("id = ?", id).Select()
	return question, err
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
