package dto

type SearchFreelances struct {
	Services []SearchFreelance `json:"services"`
	Cursor   *string           `json:"cursor"`
	HasMore  bool              `json:"hasMore"`
}
