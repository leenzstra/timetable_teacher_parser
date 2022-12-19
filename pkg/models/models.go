package models

import (
	"database/sql"
)

type Teacher struct {
	Id         int
	FIO        string
	Position   string
	Department sql.NullString
	ImageUrl   string
}

func (t Teacher) TableName() string {
	return "teachers"
}
