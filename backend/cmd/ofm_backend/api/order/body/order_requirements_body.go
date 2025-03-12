package body

import "mime/multipart"

type OrderRequirementsBody struct {
	Answers         []OrderQuestionsAnswer  `json:"answers"`
	CustomerMessage string                  `json:"customerMessage"`
	Files           []*multipart.FileHeader `json:"files"`
	OrderId         int                     `json:"orderId"`
}
