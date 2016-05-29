package conutil

//ContainerMetrics ...
type ContainerMetrics struct {
	CPUTime      uint64 `json:"CPUTime"`
	ProcessCount int    `json:"pCount"`
	TaskCount    int    `json:"tsks"`
	UsedMemory   int    `json:"usedMemory(MB)"`
	RootSize     int    `json:"rootSize(MB)"`
}
