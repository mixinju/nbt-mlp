package schedule

import (
	"fmt"
	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NewPod 创建新的容器
func NewPod(pod *core.Pod) {
	fmt.Println(pod)
	err := meta.AddMetaToScheme(nil)
	if err != nil {
		return
	}
}

func GeneratePodImage(vGpu int, vGpuMem int, name string) *core.Pod {
	panic(nil)
}

func DeletePod(podId string) {}
