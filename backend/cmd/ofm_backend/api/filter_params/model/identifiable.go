package model

type Identifiable interface {
	GetID() int64
	GetName() string
}