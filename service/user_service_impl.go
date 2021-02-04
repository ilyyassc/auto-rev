package service

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"auto-rev/config"
	"auto-rev/dao"
	"auto-rev/model"

	"fmt"
)

var userDao dao.UserDao = dao.UserDaoImpl{}

type UserServiceImpl struct{}

func (UserServiceImpl) CreateUser(user *model.Users) (u *model.Users, e error) {
	defer config.CatchError(&e)
	result, err := bcrypt.GenerateFromPassword([]byte(user.Pwd), 4)
	if err == nil {
		user.Pwd = string(result)
		// var time = model.Timestamp(time.Now())
		user.CreatedDate = time.Now()
		return userDao.CreateUser(user)
	}
	return nil, err
}

func (UserServiceImpl) GetUserById(id string) (u model.Users, e error) {
	defer config.CatchError(&e)
	u, err := userDao.GetUserById(id)
	u.Pwd = "" //do not show password !!!
	return u, err
}

func (UserServiceImpl) Login(username string, pwd string) (u model.Users, e error) {
	defer config.CatchError(&e)
	result, err := userDao.GetUserByUsername(username)
	fmt.Println("login1")
	fmt.Println(result.Id)
	fmt.Println("login2")
	if err == nil && result.Count > 0 {
		var err = bcrypt.CompareHashAndPassword([]byte(result.Pwd), []byte(pwd))
		if err == nil {
			result.Pwd = "" //do not show password !!!
			v, _ := config.GenerateToken(username, result.Id)
			result.Token = v
			return result, nil
		}
	}
	return model.Users{}, errors.New("Username/Password incorect")
}
