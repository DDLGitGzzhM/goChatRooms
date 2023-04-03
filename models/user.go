package models

import "time"

type User struct {
	ID       int
	Name     string
	Password string
	IsDelete string
	CTime    time.Time
	MTime    time.Time
}
