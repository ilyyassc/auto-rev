package dao

import (
	"auto-rev/config"
	"auto-rev/model"
)

type UserDaoImpl struct{}

func (UserDaoImpl) CreateUser(user *model.Users) (u *model.Users, e error) {
	defer config.CatchError(&e)
	result := g.Create(user)
	if result.Error == nil {
		return user, nil
	}
	return nil, result.Error
}

func (UserDaoImpl) GetUserById(id string) (u model.Users, e error) {
	defer config.CatchError(&e)
	data := model.Users{BaseModel: model.BaseModel{Id: id}}
	result := g.First(&data)

	if result.Error == nil {
		data.Count = result.RowsAffected
		return data, nil
	}
	return data, result.Error
}

func (UserDaoImpl) GetUserByUsername(username string) (u model.Users, e error) {
	defer config.CatchError(&e)
	var data model.Users
	result := g.Where("email = ?", username).Find(&data)

	if result.Error == nil {
		data.Count = result.RowsAffected
		return data, nil
	}
	return data, result.Error
}
