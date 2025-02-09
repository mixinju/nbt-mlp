package api

import (
    "net/http"
    "os"
    "path/filepath"

    "go.uber.org/zap"
    "nbt-mlp/Infrastructure/authorization"
    "nbt-mlp/application"
    "nbt-mlp/common/util"
    "nbt-mlp/common/util/errno"
    "nbt-mlp/service"

    "github.com/gin-gonic/gin"
)

var logger, _ = zap.NewProduction()

type User struct {
    auth authorization.AuthIface
    ua   application.UserAppIface
}

func (u *User) ModifyPassword(c *gin.Context) {
    // 基础校验
    var param struct {
        NewPassword string `form:"newPassword" binding:"required"`
        OldPassword string `form:"oldPassword" binding:"required"`
        UserID      string `form:"userId" binding:"required"`
    }

    if err := c.ShouldBind(&param); err != nil {
        c.JSON(http.StatusBadRequest, Error(errno.ErrBind))
        return
    }

    if !util.PasswordValid(param.NewPassword) || !util.PasswordValid(param.OldPassword) {
        c.JSON(http.StatusOK, Error(errno.ErrValidateFail))
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
        c.JSON(http.StatusBadRequest, Error(errno.ErrFileUpload))
        return
    }

    // Save the uploaded file to a temporary location
    filePath := filepath.Join("/tmp", file.Filename)
    if err := c.SaveUploadedFile(file, filePath); err != nil {
        c.JSON(http.StatusInternalServerError, Error(errno.ErrFileSave))
        return
    }

    userList, e := u.ua.ReadFromExcelFile(filePath)
    if e != nil {
        c.JSON(http.StatusOK, Error(errno.ErrFileContentEmpty))
        return
    }

    // 这里不再抛出异常,使用邮箱异步通知方式
    service.BatchRegister(userList)

    // 移除上传的文件
    if err := os.Remove(filePath); err != nil {
        c.JSON(http.StatusOK, Error(errno.InternalServerError))
        return
    }

    c.JSON(http.StatusOK, Success("文件上传成功! 正在批量注册中~"))
}

// Login 统一登录入口,通过token 区分用户身份
func (u *User) Login(c *gin.Context) {
    var param struct {
        UserID   uint64 `form:"userId" binding:"required"`
        Password string `form:"password" binding:"required"`
    }

    if err := c.ShouldBind(&param); err != nil {
        c.JSON(http.StatusBadRequest, Error(errno.ErrBind))
        return
    }

    logger.Info("api->v1->Login->", zap.Uint64("username", param.UserID))

    if !util.PasswordValid(param.Password) {
        c.JSON(http.StatusOK, ErrorWithMessage(errno.ErrValidateFail, "用户名或密码长度不符合要求"))
        return
    }

    _, err := u.ua.QueryUserByIdAndPassword(param.UserID, param.Password)
    if err != nil {
        c.JSON(http.StatusOK, Error(errno.ErrPasswordNotMatch))
    }

    token, err := u.auth.SetUpToken(param.UserID)
    if err != nil {
        c.JSON(http.StatusOK, Error(errno.ErrTokenSetUpFail))
        return
    }

    c.JSON(http.StatusOK, Success(Login{UserID: param.UserID, Token: token}))
}
