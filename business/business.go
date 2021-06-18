package business

import (
	"github.com/aljorhythm/sapere-server/model"
)

type BusinessData interface {
	GetUser(id int64) (model.User, error)
	GetReadingListByUser(user model.User) (model.ReadingList, error)
}

type RealBusiness struct {
	Data BusinessData
}
