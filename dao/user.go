package dao

import (
	"sync"

	"go.uber.org/zap"
	"nbt-mlp/dao/model"
)

type UserDao_ struct{}

var (
	instance *UserDao_
	once     sync.Once
)

func GetUserDaoInstance() *UserDao_ {
	once.Do(func() {
		instance = &UserDao_{}
	})
	return instance
}
func (u *UserDao_) QueryById(map[string]string) UserDao_ {
	DB.Model(&UserDao_{}).Where(u).First(u)
	return *u
}

func (u *UserDao_) QueryAll() []UserDao_ {
	var us []UserDao_
	DB.Model(&UserDao_{}).
		Find(&us)

	return us
}

func (_ *UserDao_) Create(u model.User) bool {
	result := DB.Model(&model.User{}).Create(&u)
	if result.Error != nil {
		logger.Error("Create user failed", zap.Error(result.Error))
		return false
	}

	return result.RowsAffected == 1
}

func (_ *UserDao_) BatchCreate(us []*model.User) bool {
	result := DB.Model(&UserDao_{}).Create(us)
	if result.Error != nil {
		logger.Error("Create user failed", zap.Error(result.Error))
		return false
	}
	return result.RowsAffected == int64(len(us))
}

func (_ *UserDao_) Update() {

}

func (_ *UserDao_) Delete() {

}
