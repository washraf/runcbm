package benchmark

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/washraf/runcbm/common"
	"github.com/washraf/runcbm/containers/config"
	"github.com/washraf/runcbm/containers/conutil"
)

//Run ..
func Run(containerID string, n int) error {
	condir, err := config.FindContainerBundle(containerID)
	if err != nil {
		return err
	}
	measuresList := make(Measures, 0)
	for i := 1; i <= n; i++ {
		measure := Measure{}
		u, err := conutil.GetContainerUtilization(containerID)
		if err != nil {
			return err
		}
		measure.ID = i
		measure.ProcessCount = u.ProcessCount
		measure.MemorySize = u.UsedMemory
		command := exec.Command("time", "-f", "%e", "runc", "checkpoint", containerID)
		command.Dir = condir
		r, err := command.CombinedOutput()
		if err != nil {
			return err
		}
		measure.CheckpointTime, _ = strconv.ParseFloat(strings.TrimSpace(string(r)), 64)
		s, err := common.FindDiskSizeMB(condir + "/checkpoint/")
		if err != nil {
			fmt.Println(condir)
			return err
		}
		measure.Checkpointsize = s
		command = exec.Command("time", "-f", "%e", "runc", "restore", "-d", containerID)
		//command.Dir = "/containers/"+container+"/"
		command.Dir = condir

		r, err = command.CombinedOutput()
		if err != nil {
			return err
		}
		measure.Restoretime, _ = strconv.ParseFloat(strings.TrimSpace(string(r)), 64)
		measuresList = append(measuresList, measure)
		//time.Sleep(time.Second * 5)
	}
	printlist(measuresList)
	return nil
}

func printlist(measuresList Measures) {
	fmt.Printf("ID\tProcessCount\tMemorySize\tCheckpointTime\tCheckpointsize\tRestoretime\n")
	for _, m := range measuresList {
		fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\n", m.ID, m.ProcessCount, m.MemorySize, m.CheckpointTime, m.Checkpointsize, m.Restoretime)
	}
}
