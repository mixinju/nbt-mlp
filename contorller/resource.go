package contorller

import "github.com/gin-gonic/gin"

type Resource struct{}

// CreatePod 创建资源
func (r *Resource) CreatePod(c *gin.Context) {
    // 验证创建次数, 避免频繁创建和删除
    // 当前用户是否已经存在运行实例,运行创建多个[数量可配置,统一配置 & 白名单配置]

    // 是否存在已经创建,但是挂起的实例

    //
}

// DeletePod 删除当前节点
func (r *Resource) DeletePod(c *gin.Context) {

    // 删除操作记录到本地缓存

    // 避免恶意执行频繁删除和创建
}

// QueryPod 查询用户名下已有的容器实例
func (r *Resource) QueryPod(c *gin.Context) {

}
