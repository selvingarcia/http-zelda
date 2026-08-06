package models

import (
	"database/sql"
)

type StatusModel struct {
	DB *sql.DB
}

type status struct{
	Code int
	Phrase string
	Title string
	Content string
	Meaning string
}

