package entity

type GPUType int

const (
    RTX3060 GPUType = iota
    RTX3090 GPUType = iota
    RTX4090 GPUType = iota
    A100    GPUType = iota
    A800    GPUType = iota
)
