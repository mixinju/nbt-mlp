package util

import (
    "strconv"

    "golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the given password using bcrypt.
func HashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}

// ComparePassword compares a hashed password with a plain text password.
func ComparePassword(hashedPassword, password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    return err == nil
}

func GetUintUserId(u string) uint64 {
    id, err := strconv.ParseUint(u, 10, 64)
    if err != nil {
        return 0
    }
    return id
}

func PasswordValid(password string) bool {
    if len(password) <= 10 || len(password) >= 20 {
        return false
    }
    return true
}
