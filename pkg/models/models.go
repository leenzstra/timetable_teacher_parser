package models

import (
	"database/sql"
)

type Teacher struct {
	Id        int
	FIO        string
	Position   string
	Department sql.NullString
}

func (t Teacher) TableName() string {
	return "teachers"
}

type TeacherEvaluation struct {
	Id int
	TeacherId int
	Mark      int
	Comment   sql.NullString
}
