package benchmark

//Measure ...
type Measure struct {
	ID             int
	ProcessCount   int
	MemorySize     int
	CheckpointTime float64
	Checkpointsize int
	Restoretime    float64
}

//Measures
type Measures []Measure
