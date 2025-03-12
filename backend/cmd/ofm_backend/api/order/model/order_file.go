package model

type OrderFile struct {
	OrderId int `db:"order_id"`
	FileId  int `db:"file_id"`
}
