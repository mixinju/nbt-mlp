package application

import (
    "context"
    "fmt"

    core "k8s.io/api/core/v1"
    meta "k8s.io/apimachinery/pkg/apis/meta/v1"
    "nbt-mlp/Infrastructure/config"
    "nbt-mlp/domain/entity"

    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/clientcmd"
)

type ContainerAppIface interface {
    DeletePod(c entity.Container) error
    CreatePod(c entity.Container) error
}

type PodApp struct {
    clientSet *kubernetes.Clientset
}

var _ ContainerAppIface = &PodApp{}

func NewPodApp() *PodApp {
    config.K8ConfigPath = "/Users/mixinju/.kube/config"
    k8sConfig, err := clientcmd.BuildConfigFromFlags("", config.K8ConfigPath)
    if err != nil {
        panic("读取k8s配置文件失败")
    }
    clientSet, err := kubernetes.NewForConfig(k8sConfig)
    if err != nil {
        log.Panic("创建pod失败")
    }
    return &PodApp{clientSet: clientSet}
}

func (p *PodApp) DeletePod(c entity.Container) error {
    err := p.clientSet.
        CoreV1().
        Pods(c.NameSpace).
        Delete(context.Background(), c.PodName, meta.DeleteOptions{})
    if err != nil {
        return fmt.Errorf("删除Pod失败: %v", err)
    }
    return nil
}

func (p *PodApp) CreatePod(c entity.Container) error {
    pod := &core.Pod{
        ObjectMeta: c.ObjectMeta(),
        Spec:       c.PodSpec(),
    }

    _, err := p.clientSet.
        CoreV1().
        Pods(c.NameSpace).
        Create(context.Background(), pod, meta.CreateOptions{})
    if err != nil {
        return fmt.Errorf("创建Pod失败: %v", err)
    }
    return nil
}
