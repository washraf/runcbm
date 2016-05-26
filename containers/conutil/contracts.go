package conutil

//ContainerMetrics ...
type ContainerMetrics struct {
	CPUTime      uint64 `json:"CPUTime"`
	ProcessCount int    `json:"pCount"`
	UsedMemory   int    `json:"usedMemory(MB)"`
	RootSize     int    `json:"rootSize(MB)"`
}
