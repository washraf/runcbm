package conutil


//helper https://blog.docker.com/2013/10/gathering-lxc-docker-containers-metrics/

//ContainerMetrics ...
type ContainerMetrics struct{
    CPUTime uint64 `json:"CPUTime"`
    UsedMemory int `json:"usedMemory(MB)"`
    RootSize uint64 `json:"rootsize"`
}