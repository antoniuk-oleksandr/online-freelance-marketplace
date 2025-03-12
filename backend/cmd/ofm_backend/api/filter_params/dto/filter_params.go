package dto

type FilterParams struct {
	Languages []FilterItem `json:"language"`
	Categories []FilterItem `json:"category" `
	Skills    []FilterItem `json:"skill"`
}
