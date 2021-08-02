package inmemory

import "github.com/aljorhythm/sapere-server/model"

type DataAccess struct{}

type ReadingList struct{}

type Reading struct{}

func (i Reading) GetUserId() int64 {
	return 1
}

type Tag struct {
	title string
}

func (t Tag) GetTitle() string {
	return t.title
}

func (i Reading) GetTags() []model.Tag {
	return []model.Tag{
		Tag{},
	}
}

type Resource struct {
	id  int64
	url string
}

func (r Resource) GetId() int64 {
	return r.id
}

func (r Resource) GetUrl() string {
	return r.url
}

func (r Resource) GetCanonicalResourceId() int64 {
	return 1
}

func (i Reading) GetResource() model.Resource {
	return Resource{}
}

func (i ReadingList) GetList() []model.Reading {
	return []model.Reading{
		Reading{},
	}
}

type User struct {
	id   int64
	name string
}

func (u User) GetId() int64 {
	return u.id
}

func (u User) GetName() string {
	return u.name
}

func (i DataAccess) GetUser(id int64) (model.User, error) {
	users := map[int64]User{
		1: {id: id, name: "John"},
	}
	if user, ok := users[id]; ok {
		return user, nil
	}
	return User{}, nil
}

func (i DataAccess) GetReadingListByUser(user model.User) (model.ReadingList, error) {
	return ReadingList{}, nil
}
