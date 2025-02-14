package application

import (
    "context"
    "fmt"

    corev1 "k8s.io/api/core/v1"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/clientcmd"
    "nbt-mlp/Infrastructure/config"
)

type ContainerAppIface interface {
    DeletePod(podId string)
}

type PodApp struct {
    clientSet *kubernetes.Clientset
}

func NewPodApp() *PodApp {
    k8sConfig, err := clientcmd.BuildConfigFromFlags("", config.K8ConfigPath)
    if err != nil {
        panic("读取k8s配置文件失败")
    }
    clientSet, err := kubernetes.NewForConfig(k8sConfig)
    return &PodApp{clientSet: clientSet}
}

func (p *PodApp) DeletePod(podId string) {
    // TODO: implement pod deletion
}

func (p *PodApp) CreatePod(image string) error {
    pod := &corev1.Pod{
        ObjectMeta: metav1.ObjectMeta{
            Name:      fmt.Sprintf("pod-%s", "sdafds"),
            Namespace: "default",
        },
        Spec: corev1.PodSpec{
            Containers: []corev1.Container{
                {
                    Name:  "main",
                    Image: image,
                },
            },
            RestartPolicy: corev1.RestartPolicyNever,
        },
    }

    _, err := p.clientSet.CoreV1().Pods("default").Create(context.Background(), pod, metav1.CreateOptions{})
    if err != nil {
        return fmt.Errorf("创建Pod失败: %v", err)
    }
    return nil
}
