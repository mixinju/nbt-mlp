package application

import "nbt-mlp/domain/entity"

type UserAppIface interface {
    Query(id uint64) (entity.User, error)
    Delete(id uint64) error
    QueryUsers(id []uint64) ([]entity.User, error)
    Update(u entity.User) error
    Save(u entity.User) error
    BatchSave(us []entity.User) error
    ReadFromFile(path string) ([]entity.User, error)
}
