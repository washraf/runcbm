package benchmark

//Measure ...
type Measure struct {
	ID                int
	ProcessCount      int
	TaskCount         int
	TotalMemorySize   int
	InRAMSize         int
	SwappedMemorySize int
	CheckpointTime    float64
	Checkpointsize    int
	Restoretime       float64
	CopyFromTime      float64
	CopyTOTime        float64
}

//Measures ...
type Measures []Measure
