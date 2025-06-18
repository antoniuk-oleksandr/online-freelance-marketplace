package dto

type OrderOverviewService struct {
	Image   string                      `json:"image"`
	Title   string                      `json:"title"`
	Package OrderOverviewServicePackage `json:"package"`
}
