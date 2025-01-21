package persistence

import (
    "sync"

    "go.uber.org/zap"
    "nbt-mlp/domain/entity"
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
func (u *UserDao_) QueryById(condition map[string]interface{}, fields []string) entity.User {
    DB.Model(&UserDao_{}).
        Select(fields).
        Where(condition).
        First(u)
    return entity.User{}
}

func (u *UserDao_) QueryAll() []UserDao_ {
    var us []UserDao_
    DB.Model(&UserDao_{}).
        Find(&us)

    return us
}

func (_ *UserDao_) Create(u entity.User) bool {
    result := DB.Model(&entity.User{}).Create(&u)
    if result.Error != nil {
        logger.Error("Create user failed", zap.Error(result.Error))
        return false
    }

    return result.RowsAffected == 1
}

func (_ *UserDao_) BatchCreate(us []*entity.User) bool {
    result := DB.Model(&UserDao_{}).Create(us)
    if result.Error != nil {
        logger.Error("Create user failed", zap.Error(result.Error))
        return false
    }
    return result.RowsAffected == int64(len(us))
}

func (_ *UserDao_) Update(u entity.User) bool {
    return false
}

func (_ *UserDao_) Delete() {

}
