package models

import "time"

type Role struct {
	Id       int64     `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Password string    `json:"password,omitempty"`
	RegDate  time.Time `json:"reg_date,omitempty"`
}
