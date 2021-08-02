package model

//CanonicalResource represents canonical resource
//different Resource might be pointing to the same CanonicalResource

type ResourceType string

type CanonicalResource interface {
	GetId() int64
	GetTitle() string
	GetUrl() string
	GetType() ResourceType
}
