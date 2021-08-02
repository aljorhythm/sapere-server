package business

import (
	"github.com/aljorhythm/sapere-server/model"
)

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func (u User) GetId() int64 {
	return u.Id
}

func (b RealBusiness) GetCurrentUser() (model.User, error) {
	return b.Data.GetUser(1234)
}

func (b RealBusiness) GetReadingFeed() (Feed, error) {
	user, error := b.GetCurrentUser()
	if error != nil {
		return Feed{}, error
	}
	list, error := b.Data.GetReadingListByUser(user)
	if error != nil {
		return Feed{}, error
	}
	return presentFeed(list), nil
}

func (b RealBusiness) GetUser(i int64) (User, error) {
	user, error := b.Data.GetUser(i)
	if error != nil {
		return User{}, error
	}
	return presentUser(user), nil
}

func presentUser(user model.User) User {
	return User{
		Id:   user.GetId(),
		Name: user.GetName(),
	}
}
