package model

import (
	"time"
)

// Host 物理主机
// 每次执行服务初始化时自动注册
type Host struct {
	// 主机ID
	Id uint64 `gorm:"primary_key,bigint,auto_increment"`

	// 网络IP
	IP string `gorm:"type:varchar(15);not null"`

	// 目前运行的容器的数量
	CountOfContainer uint8 `gorm:"default:0"`

	// GPU显卡类型
	GPUName string `gorm:"type:varchar(20);not null"`

	// 物理内存大小
	RAM uint64 `gorm:"default:0;not null"`

	// GPU内存大小
	GPUMemory uint64 `gorm:"default:0;not null"`

	// 硬盘大小
	DiskMemory uint64 `gorm:"default:0;not null"`

	// 当前物理机上运行的容器
	Containers []Container `gorm:"foreignKey:HostId"`
}

type User struct {
	// 用户主键
	Id uint64 `json:"id,omitempty" gorm:"primary_key,bigint"`

	// 用户姓名-默认使用拼音
	Name string `json:"name,omitempty" gorm:"type:varchar(20)"`

	// 年级-默认20开始,仅仅记录 23,24,25等,表示入学年份
	Grade uint8 `json:"grade,omitempty" gorm:"default:20"`

	// 权限码
	Access uint8 `json:"access,omitempty" gorm:"default:0"`

	// 用户密码
	Password string `json:"password,omitempty" gorm:"type:varchar(40);not null"`

	// 班级名称-可选
	ClassName string `json:"class_name,omitempty" gorm:"type:varchar(20)"`

	// 最后一次登录时间
	LastLoginAt time.Time `json:"last_login_at" gorm:"default:null"`

	// 用户创建时间
	CreatedAt time.Time `json:"created_at" gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP;not null"`

	// 用户组-目前仅仅支持单用户组
	AccessGroup AccessGroup `json:"access_group" gorm:"foreignKey:UserId"`

	// 一个用户可以有多个容器
	Containers []Container `json:"containers,omitempty" gorm:"foreignKey:UserId"`
}

// AccessGroup 用户组
type AccessGroup struct {
	Id     uint64 `gorm:"bigint"`
	UserId uint64 `gorm:"primary_key,bigint"`
	Name   string `gorm:"type:varchar(20)"`
}

// Container 运行的容器
type Container struct {
	Id uint64 `gorm:"primary_key,bigint"`

	// k8s集群中的 podId
	PodName string `gorm:"type:varchar(20)"`

	// 一个容器属于一个物理机器
	HostId uint64 `gorm:"bigint"`

	// 一个容器属于一个用户
	UserId uint64 `gorm:"bigint"`

	// 持久化目录挂载地址
	PersistencePath string `gorm:"type:varchar(50)"`

	// vGPU 数量
	VGPU int `gorm:"default:0"`

	// 虚拟内存大小
	VMemory uint64 `gorm:"default:0"`

	// 容器运行状态
	Status int `gorm:"default:0"`
}

type LogRecord struct {
	Id         uint64    `gorm:"primary_key,bigint"`
	Type       uint8     `gorm:"default:0"`
	Info       string    `gorm:"type:varchar(100)"`
	CreatedAt  time.Time `gorm:"default:current_timestamp;not null"`
	OperatorId uint64    `gorm:"bigint;not null"`
}
