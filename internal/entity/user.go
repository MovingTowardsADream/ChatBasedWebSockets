package entity

import "time"

type User struct {
	Id        string        `json:"id"`
	Email     string        `json:"email"`
	Username  string        `json:"login"`
	Password  string        `json:"password"`
	CreatedAt time.Duration `json:"created_at"`
}
