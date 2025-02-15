package api

import (
    "net/http"
    "os"
    "path/filepath"

    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "nbt-mlp/Infrastructure/authorization"
    "nbt-mlp/application"
    "nbt-mlp/common/util"
    "nbt-mlp/common/util/errno"
)

var logger, _ = zap.NewProduction()

type User struct {
    auth authorization.AuthIface
    ua   application.UserAppIface
    ca   application.ContainerAppIface
}

func NewUser() *User {
    return &User{
        ua:   application.NewUserAppImpl(),
        auth: authorization.NewAuthImpl(),
        ca:   application.NewPodApp(),
    }
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
    u.ua.BatchSave(userList)

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

    // 明文密码
    _, err := u.ua.QueryUserByIdAndPassword(param.UserID, param.Password)
    if err != nil {
        c.JSON(http.StatusOK, Error(err))
    }

    token, err := u.auth.SetUpToken(param.UserID)
    if err != nil {
        c.JSON(http.StatusOK, Error(errno.ErrTokenSetUpFail))
        return
    }

    c.JSON(http.StatusOK, Success(Login{UserID: param.UserID, Token: token}))
}

func (u *User) CreatePod(c *gin.Context) {
    // 基础校验

    // 忽略错误处理,除非中间件失效
    userId, _ := util.GetUserId(c)

    // 创建条件校验

    ut, err := u.ua.Query(userId)
    if err != nil {
        c.JSON(http.StatusOK, gin.H{})
        return
    }

    // 分配计算资源

    // 使用PodApp创建对应的节点

    // 保存节点并保存到数据库

    // 输出日志
}

func (u *User) DeletePod(c *gin.Context) {
    // user 基本校验

    // 容器拥有者校验,一般用户只能删除自己创建的容器,管理员删除在管理员处理逻辑中

    // 调用PodApp删除容器
}

func (u *User) QueryPod(c *gin.Context) {
    // 查询名下所有的Pod

    //
}
