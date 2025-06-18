package model

import "mime/multipart"

type OrderDeliveryBody struct {
	Message string               `json:"message"`
	Files   []*multipart.FileHeader `json:"files,omitempty"`
}
