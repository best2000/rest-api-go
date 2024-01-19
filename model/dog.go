package model

import "time"

type Dog struct {
	Id        uint32    `json:"id"`
	Name   string    `json:"name"`
	Breed   string    `json:"breed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
