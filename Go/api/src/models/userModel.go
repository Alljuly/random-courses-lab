package models

import "time"

type UserModel struct {
	ID        uint   `json:id,omitempty`
	Name      string `json:name,omitempty`
	Nick      string `json:nick,omitempty`
	Email     string `json:email,omitempty`
	Password  string `json:password,omitempty`
	CreatedAt time.Time `json:CreatedAt,omitempty`
}

