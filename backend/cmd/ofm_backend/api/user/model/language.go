package model

type Language struct {
	ID   int64    `json:"language_id" db:"language_id"`
	Name string `json:"name" db:"name"`
}

func (lang Language) GetID() int64 { return lang.ID }

func (lang Language) GetName() string { return lang.Name }
