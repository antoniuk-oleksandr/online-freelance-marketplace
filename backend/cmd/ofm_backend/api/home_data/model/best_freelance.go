package model

type BestFreelancer struct {
	Id                int64   `json:"id" db:"id"`
	FirstName         string  `json:"first_name" db:"first_name"`
	Surname           string  `json:"surname" db:"surname"`
	Rating            float64 `json:"rating" db:"rating"`
	CompletedProjects int     `json:"completed_projects" db:"completed_projects"`
	Avatar            *string `json:"avatar" db:"avatar"`
}
