package main

import (
	"net/http"
	"encoding/json"
)


type Middleware func(http.HandlerFunc) http.HandlerFunc

type MetaData interface{}

type User struct{
	 Name string `json:"name"`
	 Email string `json:"email"`
	 Phone string `json:"phone"`
}

func (u *User ) ToJson() ([]byte, error) {
	return json.Marshal(u)
}