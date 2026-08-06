package models

import (
	"database/sql"
	"errors"
)

var NoRows = errors.New("sql: no rows found in the database")

type StatusModel struct {
	DB *sql.DB
}

type Status struct{
	Code int
	Phrase string
	Title string
	Content string
	Meaning string
}

func (s *StatusModel) GetStatus(n int)(Status, error){
	query := `select code, phrase , title, content, meaning from codes where code = $1;`

	t := Status{}

	row := s.DB.QueryRow(query,n)

	err := row.Scan(&t.Code, &t.Phrase, &t.Title, &t.Content, &t.Meaning)
	if err!=nil{
		if errors.Is(err, sql.ErrNoRows){
			return Status{}, NoRows
		}else{
			return Status{}, err
		}
		return Status{}, err
	}
	return t,nil

}
