package persistence

import (
    "gorm.io/gorm"
    "nbt-mlp/domain/entity"
    "nbt-mlp/domain/repository"
)

type UserDao struct {
    db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
    return &UserDao{db: db}
}

var _ repository.UserRepositoryIface = &UserDao{}

func (ud *UserDao) Save(u entity.User) error {
    //TODO implement me
    panic("implement me")
}

func (ud *UserDao) BatchSave(us []entity.User) error {
    //TODO implement me
    panic("implement me")
}

func (ud *UserDao) Query(id uint64) (entity.User, error) {
    //TODO implement me
    panic("implement me")
}

func (ud *UserDao) QueryUsers(id []uint64) ([]entity.User, error) {
    //TODO implement me
    panic("implement me")
}

func (ud *UserDao) Update(u entity.User) error {
    //TODO implement me
    panic("implement me")
}
