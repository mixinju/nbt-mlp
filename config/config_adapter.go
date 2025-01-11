package config

// ResourceNeverRecycleWhiteList 资源回收白名单
func ResourceNeverRecycleWhiteList() []uint64 {
	panic("never recycle white list")
}

func CheckInRecycleBlackList(userId uint64) bool {
	panic("recycle black list")
}
