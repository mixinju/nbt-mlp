package dao

import (
	"nbt-mlp/dao/model"
	"testing"
)

func TestBatchRegisterUser(t *testing.T) {

	user := make([]*model.User, 0)
	user = append(user, &model.User{ID: 1, Name: "张三", ClassName: "计算机206"})
	user = append(user, &model.User{ID: 2, Name: "李四", ClassName: "计算机206"})
	user = append(user, &model.User{ID: 3, Name: "王五", ClassName: "计算机206"})

	// 1. Create a user dao instance

	GetUserDaoInstance().BatchCreate(user)
}
