package model

type Tag interface {
	GetTitle() string
}

type Reading interface {
	GetUserId() int64
	GetTags() []Tag
	GetResource() Resource
}
