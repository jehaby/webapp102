package entity

import "time"

type User struct {
	Login     string
	Email     string
	Password  string
	LastLogin time.Time
}
