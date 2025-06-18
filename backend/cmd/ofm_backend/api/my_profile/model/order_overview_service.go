package model

type OrderOverviewService struct {
	Image   string                      `json:"image" db:"image"`
	Title   string                      `json:"title" db:"title"`
	Package OrderOverviewServicePackage `json:"package" db:"package"`
}
