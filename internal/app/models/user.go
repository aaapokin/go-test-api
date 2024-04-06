package models

type Token struct{
	Token string `json:"token"`
}

type User struct{
	ID string `json:"id"`
	Login string `json:"login"`
	Password string `json:"password"`
}