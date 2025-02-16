package application

import (
    "fmt"
    "testing"

    "nbt-mlp/domain/entity"
)

func TestBatchSave(t *testing.T) {
    users := []entity.User{
        {ID: 0210001, Name: "User1", ClassName: "Class1"},
        {ID: 0110002, Name: "User2", ClassName: "Class2"},
    }

    us := NewUserAppImpl()

    us.BatchSave(users)

}

func TestSave(t *testing.T) {
    user := entity.User{
        ID: 3200421039, Name: "米新举", ClassName: "计算机206",
    }

    us := NewUserAppImpl()

    us.Save(user)
}

func TestLogin(t *testing.T) {

    us := NewUserAppImpl()
    u, err := us.QueryUserByIdAndPassword(210001, generateDefaultPassword(210001, "扈书豪"))
    if err != nil {
        fmt.Println("Error: ", err)
    }

    fmt.Println("User: ", u)
}
