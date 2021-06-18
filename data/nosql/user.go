package nosql

import "github.com/aljorhythm/sapere-server/model"

type DbUser struct {
	id   int64
	name string
}

func (d DbUser) GetId() int64 {
	return d.id
}

func (d DbUser) GetName() string {
	return d.name
}

func (DataAccess) GetUser(id int64) (model.User, error) {
	return DbUser{}, nil
}
