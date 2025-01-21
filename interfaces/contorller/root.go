package contorller

import "fmt"

// 后门接口

// InitUserPassword 管理员使用,一键初始化密码
func InitUserPassword(id uint64) {

	fmt.Println("InitUserPassword")
}
