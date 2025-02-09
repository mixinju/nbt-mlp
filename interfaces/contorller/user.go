package contorller

import (
    "fmt"
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
    auth authorization.AuthInterface
    ua   application.UserAppIface
}

func (u *User) ModifyPassword(c *gin.Context) {
    // 基础校验
    newPwd := c.PostForm("newPassword")
    oldPwd := c.PostForm("oldPassword")
    userId := c.PostForm("userId")
    if !util.PasswordValid(newPwd) || !util.PasswordValid(oldPwd) {
        c.JSON(http.StatusOK, util.Error(errno.ErrValidateFail))

        return
    }
    fmt.Println("userId", userId)

    // 1. 用户存在
    // 2. 两次密码不一致
    // 3. 密码长度

    // 更新密码
}

func (u *User) BatchRegisterByExcelFile(c *gin.Context) {
    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, util.Error(errno.ErrFileUpload))
        return
    }

    // Save the uploaded file to a temporary location
    filePath := filepath.Join("/tmp", file.Filename)
    if err := c.SaveUploadedFile(file, filePath); err != nil {
        c.JSON(http.StatusInternalServerError, util.Error(errno.ErrFileSave))
        return
    }

    userList, e := util.ReadUserInfoFromExcel(filePath)
    if e != nil {
        c.JSON(http.StatusOK, util.Error(errno.ErrFileContentEmpty))
        return
    }

    // 这里不再抛出异常,使用邮箱异步通知方式
    service.BatchRegister(userList)

    // 移除上传的文件
    if err := os.Remove(filePath); err != nil {
        c.JSON(http.StatusOK, util.Error(errno.InternalServerError))
        return
    }

    c.JSON(http.StatusOK, util.Success("文件上传成功! 正在批量注册中~"))
}

// Login 统一登录入口,通过token 区分用户身份
func (u *User) Login(c *gin.Context) {
    n := c.PostForm("userId")
    p := c.PostForm("password")
    logger.Info("api->v1->Login->", zap.String("username", n))

    if !util.PasswordValid(p) {
        // 尝试使用自定义状态码
        c.JSON(http.StatusOK, util.ErrorWithMessage(errno.ErrValidateFail, "用户名或密码长度不符合要求"))
    }

    uintId := util.GetUintUserId(n)
    token, err := service.Login(uintId, p)
    if err != *errno.OK {
        c.JSON(http.StatusOK, util.Error(&err))
        return
    }

    c.JSON(http.StatusOK, util.Success(util.Login{UserID: uintId, Token: token}))
}
