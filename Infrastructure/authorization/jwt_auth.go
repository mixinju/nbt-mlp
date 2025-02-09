package authorization

import (
    "github.com/dgrijalva/jwt-go"
    "go.uber.org/zap"
)

var logger, _ = zap.NewProduction()

type AuthInterface interface {
    SetUpToken(userID int64) (string, error)
    CreateToken(claims jwt.MapClaims) (string, error)
    ParserToken(tokenString string) (*jwt.MapClaims, error)
}
