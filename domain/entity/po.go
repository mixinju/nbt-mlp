package entity

import (
    "fmt"
    "time"
)

// Host 物理主机
// 每次执行服务初始化时自动注册
type Host struct {
    ID               uint64      `gorm:"primaryKey;autoIncrement"`  // 主机ID
    IP               string      `gorm:"type:varchar(15);not null"` // 网络IP
    CountOfContainer uint8       `gorm:"default:0"`                 // 目前运行的容器的数量
    GPUName          string      `gorm:"type:varchar(20);not null"` // GPU显卡类型
    RAM              uint64      `gorm:"default:0;not null"`        // 物理内存大小
    GPUMemory        uint64      `gorm:"default:0;not null"`        // GPU内存大小
    DiskMemory       uint64      `gorm:"default:0;not null"`        // 硬盘大小
    Containers       []Container `gorm:"foreignKey:HostID"`         // 当前物理机上运行的容器
}

// User 用户
type User struct {
    ID          uint64      `json:"id,omitempty" gorm:"primaryKey"`                                      // 用户主键
    Name        string      `json:"name,omitempty" gorm:"type:varchar(20)"`                              // 用户姓名-默认使用拼音
    Grade       uint8       `json:"grade,omitempty" gorm:"default:20"`                                   // 年级-默认20开始,仅仅记录 23,24,25等,表示入学年份
    Access      uint8       `json:"access,omitempty" gorm:"default:0"`                                   // 权限码
    Password    string      `json:"password,omitempty" gorm:"type:varchar(64);not null"`                 // 用户密码
    ClassName   string      `json:"class_name,omitempty" gorm:"type:varchar(20)"`                        // 班级名称-可选
    LastLoginAt time.Time   `json:"last_login_at" gorm:"default:null"`                                   // 最后一次登录时间
    CreatedAt   time.Time   `json:"created_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP;not null"` // 用户创建时间
    AccessGroup AccessGroup `json:"access_group" gorm:"foreignKey:UserID"`                               // 用户组-目前仅仅支持单用户组
    Containers  []Container `json:"containers,omitempty" gorm:"foreignKey:UserID"`                       // 一个用户可以有多个容器
}

// AccessGroup 用户组
type AccessGroup struct {
    ID     uint64 `gorm:"primaryKey"`       // 用户组ID
    UserID uint64 `gorm:"not null"`         // 用户ID
    Name   string `gorm:"type:varchar(20)"` // 用户组名称
}

// Container 运行的容器
type Container struct {
    ID              uint64 `gorm:"primaryKey"`       // 容器ID
    PodName         string `gorm:"type:varchar(20)"` // k8s集群中的 podId
    Image           string `gorm:"type:varchar(50)"` // 容器镜像
    HostID          uint64 `gorm:"not null"`         // 一个容器属于一个物理机器
    UserID          uint64 `gorm:"not null"`         // 一个容器属于一个用户
    PersistencePath string `gorm:"type:varchar(50)"` // 持久化目录挂载地址
    VGPU            int    `gorm:"default:0"`        // vGPU 数量
    VMemory         uint64 `gorm:"default:0"`        // 虚拟内存大小
    Status          int    `gorm:"default:0"`        // 容器运行状态
}

// LogRecord 日志记录
type LogRecord struct {
    ID         uint64    `gorm:"primaryKey"`                         // 日志ID
    Type       uint8     `gorm:"default:0"`                          // 日志类型
    Info       string    `gorm:"type:varchar(100)"`                  // 日志信息
    CreatedAt  time.Time `gorm:"default:current_timestamp;not null"` // 创建时间
    OperatorID uint64    `gorm:"not null"`                           // 操作者ID
}

func (u *User) Check() error {
    if len(u.Name) == 0 || len(u.Name) > 20 {
        return fmt.Errorf("name must be between 1 and 20 characters")
    }
    if len(u.Password) == 0 || len(u.Password) > 64 {
        return fmt.Errorf("password must be between 1 and 64 characters")
    }
    if u.Grade < 20 {
        return fmt.Errorf("grade must be 20 or higher")
    }
    if len(u.ClassName) > 20 {
        return fmt.Errorf("ClassName must be 20 characters or less")
    }
    return nil
}
