package entity

type User struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
