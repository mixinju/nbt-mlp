package repository

import "nbt-mlp/domain/entity"

type UserRepositoryIface interface {
    Save(u entity.User) error
    BatchSave(us []entity.User) error
    Query(id uint64) (entity.User, error)
    QueryUsers(id []uint64) ([]entity.User, []error)
    Update(u entity.User) error
}
