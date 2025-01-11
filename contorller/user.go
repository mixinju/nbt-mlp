package contorller

import (
	"net/http"
	"os"
	"path/filepath"

	"nbt-mlp/contorller/result"
	"nbt-mlp/service"
	"nbt-mlp/util"
	"nbt-mlp/util/errno"

	"github.com/gin-gonic/gin"
)

type User struct{}

func (u *User) ModifyPassword(c *gin.Context) {
	// 基础校验
	// 1. 用户存在
	// 2. 两次密码不一致(可配置)
	// 3. 密码长度

	// 更新密码
}

func (u *User) BatchRegisterByExcelFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, result.Error(errno.ErrFileUpload))
		return
	}

	// Save the uploaded file to a temporary location
	filePath := filepath.Join("/tmp", file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, result.Error(errno.ErrFileSave))
		return
	}

	userList, e := util.ReadUserInfoFromExcel(filePath)
	if e != *errno.OK {
		c.JSON(http.StatusOK, result.Error(&e))
	}

	// 这里不再抛出异常,使用邮箱异步通知方式
	service.BatchRegister(userList)

	// 错误码定义不合理
	if err := os.Remove(filePath); err != nil {
		c.JSON(http.StatusOK, result.Error(errno.InternalServerError))
		return
	}

	c.JSON(http.StatusOK, result.Success("文件上传成功! 正在批量注册中~"))
}

// Login 统一登录入口,通过token 区分用户身份
func (u *User) Login(c *gin.Context) {
}
