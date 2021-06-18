package model

type User interface {
	GetId() int64
	GetName() string
}
