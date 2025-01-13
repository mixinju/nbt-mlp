package service

import (
	"testing"

	"nbt-mlp/dao"
	"nbt-mlp/dao/model"
)

func TestBatchRegister(t *testing.T) {
	// 1. Create a user service instance
	dao.InitDB()

	user := make([]model.User, 0)
	user = append(user, model.User{ID: 1, Name: "张三", ClassName: "计算机206"})
	user = append(user, model.User{ID: 2, Name: "李四", ClassName: "计算机206"})
	user = append(user, model.User{ID: 3, Name: "王五", ClassName: "计算机206"})

	BatchRegister(user)
}
