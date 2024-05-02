package models

type UserModels struct {
	Id       int64  `json:"user_id"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}