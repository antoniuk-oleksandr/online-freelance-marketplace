package model

import "encoding/json"

type FilterParams struct {
	Languages  json.RawMessage `db:"languages"`
	Categories json.RawMessage `db:"categories"`
	Skills     json.RawMessage `db:"skills"`
}
