package interfaces

import (
    "github.com/gin-gonic/gin"
    "nbt-mlp/interfaces/api"
)

func Router(u api.User) *gin.Engine {

    e := gin.Default()
    e.POST("/user/batch/register", u.BatchRegisterByExcelFile)
    e.POST("/user/login", u.Login)

    _ = e.Run(":8080")

    return e
}
