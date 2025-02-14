package util

import (
    "fmt"
    "testing"
)

func TestHashPassword(t *testing.T) {
    p1 := "1008611"
    //p2 := "mixinju8980"

    h1, _ := HashPassword(p1)
    h2, _ := HashPassword(p1)

    fmt.Println("Hashed Password 1: ", h1)
    fmt.Println("Hashed Password 2: ", h2)

    if !ComparePassword(h1, p1) {
        t.Error("ComparePassword failed")
    }
}
