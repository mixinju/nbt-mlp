package model

import "time"

// Host 物理主机
// 每次执行服务初始化时自动注册
type Host struct {
	// 主机ID
	Id uint64

	// 网络IP
	IP string

	// 目前运行的容器的数量
	CountOfContainer uint8

	// GPU显卡类型
	GPUName string

	// 物理内存大小
	RAM uint64

	// GPU内存大小
	GPUMemory uint64

	// 硬盘大小
	DiskMemory uint64

	// 当前物理机上运行的容器
	Containers []Container
}

type User struct {
	// 用户主键
	Id uint64

	// 用户姓名-默认使用拼音
	Name string

	// 年级-默认20开始,仅仅记录 23,24,25等,表示入学年份
	Grade uint8

	// 权限码
	Access uint8

	// 用户密码
	Password string

	// 班级名称-可选
	ClassName string

	// 最后一次登录时间
	LastLoginAt time.Time

	// 用户创建时间
	CreatedAt time.Time

	// 用户组-目前仅仅支持单用户组
	AccessGroup AccessGroup

	// 一个用户可以有多个容器
	Containers []Container
}

// AccessGroup 用户组
type AccessGroup struct {
	Id   uint64
	Name string
}

// Container 运行的容器
type Container struct {
	Id uint64

	// k8s集群中的 podId
	PodId string

	// 一个容器属于一个物理机器
	Host Host

	// 一个容器属于一个用户
	User User

	// 持久化目录挂载地址
	PersistencePath string

	// vGPU 数量
	VGPU int

	// 虚拟内存大小
	VMemory uint64

	// 容器运行状态
	Status int
}
