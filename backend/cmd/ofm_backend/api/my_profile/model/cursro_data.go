package model

import "time"

type CursorData struct {
	LastId   int       `json:"lastId"`
	LastDate time.Time `json:"lastDate"`
}
