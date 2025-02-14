package persistence

import (
	"nbt-mlp/domain/entity"
	"nbt-mlp/domain/repository"

	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao() *UserDao {
	InitDB()
	return &UserDao{DB}
}

var _ repository.UserRepositoryIface = &UserDao{}

func (ud *UserDao) Save(u entity.User) error {
	result := ud.db.Create(&u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ud *UserDao) BatchSave(us []entity.User) error {
	result := ud.db.Create(&us)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ud *UserDao) Query(id uint64) (entity.User, error) {
	var user entity.User
	// 使用gorm查询用户
	result := ud.db.First(&user, id)
	if result.Error != nil {
		return entity.User{}, result.Error
	}
	return user, nil
}

func (ud *UserDao) QueryUsers(ids []uint64) ([]entity.User, []error) {
	var users []entity.User
	var errs []error

	result := ud.db.Find(&users, ids)
	// Initialize the error slice with nil values
	errs = make([]error, len(ids))

	// If there's a general query error, set it for all entries
	if result.Error != nil {
		for i := range errs {
			errs[i] = result.Error
		}
		return users, errs
	}

	// Check if all IDs were found
	if len(users) != len(ids) {
		foundIDs := make(map[uint64]bool)
		for _, user := range users {
			foundIDs[user.ID] = true
		}

		// Add errors for missing IDs
		for _, id := range ids {
			if !foundIDs[id] {
				errs = append(errs, gorm.ErrRecordNotFound)
			}
		}
	}

	return users, errs
}

func (ud *UserDao) Update(u entity.User) error {
	result := ud.db.Save(&u)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
