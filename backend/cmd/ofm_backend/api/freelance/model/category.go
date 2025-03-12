package model

type Category struct {
	ID   int64  `json:"category_id" db:"category_id"`
	Name string `json:"name" db:"name"`
}

func (c Category) GetID() int64    { return c.ID }

func (c Category) GetName() string { return c.Name }
