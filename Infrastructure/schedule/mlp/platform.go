package mlp

import (
	"fmt"
	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 连接 k8s 集群初始化
func init() {

}

// 获取所有物理机器节点信息
// 上游传递物理节点信息
func initAllHost(pod *core.Pod) {
	group := meta.APIGroup{Name: "ok"}

	fmt.Println(group)
}
