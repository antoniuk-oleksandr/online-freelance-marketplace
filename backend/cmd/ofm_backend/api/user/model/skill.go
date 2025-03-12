package model

type Skill struct {
	ID   int64  `json:"skill_id"`
	Name string `json:"name"`
}

func (s Skill) GetID() int64 { return s.ID }

func (s Skill) GetName() string { return s.Name }
