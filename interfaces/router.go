package interfaces

import (
	"nbt-mlp/Infrastructure/middleware"
	"nbt-mlp/interfaces/api"

	"github.com/gin-gonic/gin"
)

func Router(u *api.User) *gin.Engine {

	e := gin.Default()

	// 使用跨域中间件
	e.Use(middleware.Cors())

	e.POST("/user/batch/register", u.BatchRegisterByExcelFile)
	e.POST("/user/login", u.Login)

	_ = e.Run(":8080")

	return e
}
