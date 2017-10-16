package main

import "github.com/go-pg/pg"
import "fmt"

// Question model
type Question struct {
	ID      int    `json:"id"`
	Text    string `json:"text"`
	Upvotes int    `json:"upvotes"`
	Author  string `json:"author"`
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
	fmt.Println(err)
	return err
}

func (dal *DAL) Upvote(id int) error {
	_, err := dal.conn.Model(&Question{}).Where("id = ?", id).Set("upvotes = upvotes + 1").Update()
	return err
}
