package dto

type BestFreelancer struct {
	Id                int64   `json:"id"`
	FirstName         string  `json:"firstName"`
	Surname           string  `json:"surname"`
	Rating            float64 `json:"rating"`
	CompletedProjects int     `json:"completedProjects"`
	Avatar            *string  `json:"avatar"`
}
