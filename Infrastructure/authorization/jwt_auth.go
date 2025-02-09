package authorization

import (
    "github.com/dgrijalva/jwt-go"
    "go.uber.org/zap"
    "nbt-mlp/common/util/errno"
)

var logger, _ = zap.NewProduction()

type AuthIface interface {
    SetUpToken(userID uint64) (string, *errno.Errno)
    ParserToken(tokenString string) (*jwt.MapClaims, *errno.Errno)
}
