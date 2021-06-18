package model

type Resource interface {
	GetId() int64
	GetUrl() string
	GetCanonicalResourceId() int64
}
