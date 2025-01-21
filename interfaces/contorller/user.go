package contorller

import (
    "net/http"
    "os"
    "path/filepath"

    "go.uber.org/zap"
    "nbt-mlp/contorller/result"
    "nbt-mlp/service"
    "nbt-mlp/util"
    "nbt-mlp/util/errno"

    "github.com/gin-gonic/gin"
)

var logger, _ = zap.NewProduction()

type User struct{}

func (u *User) ModifyPassword(c *gin.Context) {
    // 基础校验
    newPwd := c.PostForm("newPassword")
    oldPwd := c.PostForm("oldPassword")
    userId := c.PostForm("userId")
    if !util.PasswordValid(newPwd) || !util.PasswordValid(oldPwd) {
        c.JSON(http.StatusOK, result.Error(errno.ErrValidateFail))
        return
    }

    // 1. 用户存在
    // 2. 两次密码不一致
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

    // 移除上传的文件
    if err := os.Remove(filePath); err != nil {
        c.JSON(http.StatusOK, result.Error(errno.InternalServerError))
        return
    }

    c.JSON(http.StatusOK, result.Success("文件上传成功! 正在批量注册中~"))
}

// Login 统一登录入口,通过token 区分用户身份
func (u *User) Login(c *gin.Context) {
    n := c.PostForm("userId")
    p := c.PostForm("password")
    logger.Info("api->v1->Login->", zap.String("username", n))

    if !util.PasswordValid(p) {
        // 尝试使用自定义状态码
        c.JSON(http.StatusOK, result.ErrorWithMessage(errno.ErrValidateFail, "用户名或密码长度不符合要求"))
    }

    uintId := util.GetUintUserId(n)
    token, err := service.Login(uintId, p)
    if err != *errno.OK {
        c.JSON(http.StatusOK, result.Error(&err))
        return
    }

    c.JSON(http.StatusOK, result.Success(result.Login{UserID: uintId, Token: token}))
}
