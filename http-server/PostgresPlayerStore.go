package main

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

func NewPostgresPlayerStore() *PostgresPlayerStore {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "postgres",
		Database: "go_tdd",
	})
	return &PostgresPlayerStore{db}
}

type User struct {
	Name  string `sql:",pk"`
	Score int
}
type PostgresPlayerStore struct {
	db *pg.DB
}

func (p *PostgresPlayerStore) GetPlayerScore(name string) int {
	user := &User{Name: name}
	err := p.db.Select(user)
	if err != nil {
		user.Score = 0
		err = p.db.Insert(user)
		if err != nil {
			panic(err)
		}
		return 0
	}
	return user.Score
}

func (p *PostgresPlayerStore) RecordWin(name string) {
	user := &User{Name: name}
	err := p.db.Select(user)
	if err != nil {
		user.Score = 1
		err = p.db.Insert(user)
	} else {
		user.Score++
		err = p.db.Update(user)
	}
	if err != nil {
		panic(err)
	}
}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*User)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists: true,
			// Temp:        true,
		})
		if err != nil {
			panic(err)
		}
	}
	return nil
}
