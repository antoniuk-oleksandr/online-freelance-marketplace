package model

import "github.com/jmoiron/sqlx"

type NamedQuerier interface {
	NamedQuery(query string, arg any) (*sqlx.Rows, error)
}
