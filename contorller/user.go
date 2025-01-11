package contorller

import "github.com/gin-gonic/gin"

type User struct{}

func (u *User) ModifyPassword(c *gin.Context) {
    // 基础校验
    // 1. 用户存在
    // 2. 两次密码不一致(可配置)
    // 3. 密码长度

    // 更新密码
}

func (u *User) BatchRegisterByExcelFile(c *gin.Context) {
    // 基本文件校验,大小,格式

    // 解析为 []User 形式,异步生成注册结果
}

// Login 统一登录入口,通过token 区分用户身份
func (u *User) Login(c *gin.Context) {

}
