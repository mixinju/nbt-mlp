package authorization

import (
    "fmt"
    "time"

    "github.com/dgrijalva/jwt-go"
    "go.uber.org/zap"
    "nbt-mlp/Infrastructure/config"
    "nbt-mlp/common/util/errno"
)

var logger, _ = zap.NewProduction()

type AuthIface interface {
    SetUpToken(userID uint64) (string, *errno.Errno)
    ParserToken(tokenString string) (*jwt.MapClaims, *errno.Errno)
}

type AuthImpl struct {
}

var _ AuthIface = &AuthImpl{}

func NewAuthImpl() AuthIface {
    return &AuthImpl{}
}
func (a *AuthImpl) SetUpToken(userID uint64) (string, *errno.Errno) {
    claims := jwt.MapClaims{
        "userId":    userID,
        "createdAt": time.Now().Unix(),
        "expiredAt": time.Now().Add(time.Hour * 24 * 7).Unix(), // Token expires in 24*7 hours
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString([]byte(config.JwtKey))
    if err != nil {
        logger.Error("Failed to sign token", zap.Error(err))
        return "", errno.ErrTokenSetUpFail
    }
    return tokenString, nil
}

func (a *AuthImpl) ParserToken(tokenString string) (*jwt.MapClaims, *errno.Errno) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(config.JwtKey), nil
    })

    if err != nil {
        logger.Error("Failed to parse token", zap.Error(err))
        return nil, errno.ErrTokenInvalid
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return &claims, nil
    } else {
        return nil, errno.ErrTokenInvalid
    }
}
