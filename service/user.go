package service

import (
	"fmt"
	"nbt-mlp/dao/model"
)

func Login() {

}

func Register(u model.User) {
	// 先校验基本参数是否缺失

	// 检查对应的学号是否存在

	// 自动生成密码-默认密码:学号添加姓名拼音
}

// BatchRegister 批量注册
func BatchRegister(users []model.User) {

	// 批量注册- 使用 go协程的方式
	// 使用 channel 作为通信方式,通知注册结果
	for _, user := range users {
		go Register(user)
		fmt.Println(user)
	}

}
