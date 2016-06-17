package conutil

//ContainerMetrics ...
type ContainerMetrics struct {
	CPUTime       uint64 `json:"CPUTime"`
	ProcessCount  int    `json:"pCount"`
	TaskCount     int    `json:"tsks"`
	UsedRAM       int    `json:"usedRAM(MB)"`
	SwappedMemory int    `json:"swappedMemory(MB)"`
	TotalMemory   int    `json:"swappedMemory(MB)"`
	RootSize      int    `json:"rootSize(MB)"`
}
