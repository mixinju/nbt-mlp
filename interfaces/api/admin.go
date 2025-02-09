package api

import "fmt"

type Admin struct{}

func NewAdmin() *Admin {
    return &Admin{}
}

func (admin *Admin) Login(username, password string) (string, error) {
    fmt.Println(username, password)

    return "", nil
}

// QueryPod 查询当前所有正在运行的Pod情况
func (admin *Admin) QueryPod() error {
    return nil
}

// PausePod  强制挂起某一个正在运行的Pod
func (admin *Admin) PausePod() error {
    panic("implement me")
}

// ResumePod  恢复一个被挂起的Pod
func (admin *Admin) ResumePod() error {
    panic("implement me")
}

func (admin *Admin) DeletePod() error {
    panic("implement me")
}
