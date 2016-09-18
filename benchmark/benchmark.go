package benchmark

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/washraf/runcbm/common"
	"github.com/washraf/runcbm/containers/config"
	"github.com/washraf/runcbm/containers/conutil"
)

//Run ..
func Run(d int, containerID string, n int, move int, other string) error {
	condir, err := config.FindContainerBundle(containerID)
	if err != nil {
		return err
	}
	err = deleteCheckPointData(condir)
	if err != nil {
		return err
	}

	measuresList := make(Measures, 0)
	for i := 1; i <= n; i++ {

		fmt.Println("trial number ", i)
		//time.Sleep(time.Second * 5)
		fmt.Println("Let it RUN for 5 Seconds")
		time.Sleep(time.Second * 5)
		measure := Measure{}
		u, err := conutil.GetContainerUtilization(containerID)
		if err != nil {
			fmt.Println("Read Utilization error")
			return err
		}
		measure.ID = d * i
		measure.ProcessCount = u.ProcessCount
		measure.TaskCount = u.TaskCount
		measure.InRAMSize = u.UsedRAM
		measure.SwappedMemorySize = u.SwappedMemory
		measure.TotalMemorySize = u.TotalMemory

		//command := exec.Command("time", "-f", "%e", "runc", "checkpoint", "--tcp-established", "--empty-ns", "network", "--work-path", ".", containerID)
		command := exec.Command("runc", "checkpoint", "--tcp-established", "--empty-ns", "network", "--work-path", ".", containerID)
		command.Dir = condir
		err = command.Run()
		if err != nil {
			fmt.Println("Checkpoint error")
			return err
		}
		measure.CheckpointTime = 0
		xx, _ := common.ReadUsingCrit(condir, "stats-dump", "dump", "freezing_time", "frozen_time")
		for _, x := range xx {
			fmt.Println(x)
			t, _ := strconv.ParseFloat(strings.TrimSpace(string(x)), 64)
			fmt.Println(t / 1000000.0)
			measure.CheckpointTime += (t / 1000000.0)
		}

		//measure.CheckpointTime, _ = strconv.ParseFloat(strings.TrimSpace(string(r)), 64)
		s, err := common.FindDiskSizeMB(condir + "/checkpoint/")
		if err != nil {
			fmt.Println("Read Checkpoint Disk size error")

			fmt.Println(condir)
			return err
		}
		measure.Checkpointsize = s
		if move == 0 {
			fmt.Println("Sleep for 5 Seconds between checkpoint and restore")
			time.Sleep(time.Second * 5)
		}

		if move == 3 || move == 4 {
			command := exec.Command("time", "-f", "%e", "mv", "checkpoint/", other)
			command.Dir = condir
			r, err := command.CombinedOutput()
			if err != nil {
				fmt.Println("1st copy fail")
				return err
			}
			ct1, _ := strconv.ParseFloat(strings.TrimSpace(string(r)), 64)
			fmt.Println("sleep between move and back")
			time.Sleep(time.Second * 5)

			command = exec.Command("time", "-f", "%e", "mv", "checkpoint/", condir)
			command.Dir = other
			r, err = command.CombinedOutput()
			if err != nil {
				fmt.Println("2st copy fail")
				return err
			}
			ct2, _ := strconv.ParseFloat(strings.TrimSpace(string(r)), 64)
			measure.CopyTOTime = ct1

			if move == 3 {
				measure.CopyFromTime = ct2
			}
			err = deleteCheckPointData(other)
			if err != nil {
				fmt.Println("delete from copy err")
				return err
			}
		}

		//The Restore Process
		//command = exec.Command("time", "-f", "%e", "runc", "restore", "-d", "--tcp-established", "--work-path", ".", containerID)
		command = exec.Command("runc", "restore", "-d", "--tcp-established", "--work-path", ".", containerID)
		//command.Dir = "/containers/"+container+"/"
		command.Dir = condir

		err = command.Run()
		if err != nil {
			fmt.Println("Restore error")
			return err
		}
		xx, _ = common.ReadUsingCrit(condir, "stats-restore", "restore", "restore_time")
		measure.Restoretime = 0
		for _, x := range xx {
			fmt.Println(x)
			t, _ := strconv.ParseFloat(strings.TrimSpace(string(x)), 64)
			measure.Restoretime += (t / 1000000.0)
		}
		//measure.Restoretime, _ = strconv.ParseFloat(strings.TrimSpace(string(r)), 64)
		measuresList = append(measuresList, measure)
		err = writetoFile(logFile, measure)
		if err != nil {
			return err
		}
		err = deleteCheckPointData(condir)
		if err != nil {
			return err
		}
	}
	printlist(measuresList)
	return nil
}

func deleteCheckPointData(condir string) error {
	delcommand := exec.Command("rm")
	delcommand.Dir = condir
	delcommand.Args = append(delcommand.Args, "-rf")
	delcommand.Args = append(delcommand.Args, "checkpoint/")
	err := delcommand.Run()
	if err != nil {
		fmt.Println("Delete Checkpoint error")
		return err
	}
	return nil
}

func printlist(measuresList Measures) {
	fmt.Printf("ID\tProcessCount\tTaskCount\tMemorySize\tInRam\tSwapped\tCheckpointTime\tCheckpointsize\tRestoretime\tCopyTOTime\tCopyFrom\n")
	for _, m := range measuresList {
		fmt.Printf("%v\t\t%v\t\t%v\t\t%v\t\t%v\t\t%v\t\t%v\t\t%v\t\t%v\t%v\t%v\n", m.ID, m.ProcessCount, m.TaskCount, m.TotalMemorySize, m.InRAMSize, m.SwappedMemorySize, m.CheckpointTime, m.Checkpointsize, m.Restoretime, m.CopyTOTime, m.CopyFromTime)
	}
}
func writetoFile(filename string, m Measure) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	var buffer bytes.Buffer
	buffer.WriteString(strconv.FormatInt(int64(m.ID), 10))
	buffer.WriteString(",")
	buffer.WriteString(strconv.FormatInt(int64(m.ProcessCount), 10))
	buffer.WriteString(",")
	buffer.WriteString(strconv.FormatInt(int64(m.TaskCount), 10))
	buffer.WriteString(",")
	buffer.WriteString(strconv.FormatInt(int64(m.TotalMemorySize), 10))
	buffer.WriteString(",")
	buffer.WriteString(strconv.FormatInt(int64(m.InRAMSize), 10))
	buffer.WriteString(",")
	buffer.WriteString(strconv.FormatInt(int64(m.SwappedMemorySize), 10))
	buffer.WriteString(",")
	buffer.WriteString(floatToString(m.CheckpointTime))
	buffer.WriteString(",")
	buffer.WriteString(strconv.FormatInt(int64(m.Checkpointsize), 10))
	buffer.WriteString(",")
	buffer.WriteString(floatToString(m.Restoretime))
	buffer.WriteString(",")
	buffer.WriteString(floatToString(m.CopyTOTime))
	buffer.WriteString(",")
	buffer.WriteString(floatToString(m.CopyFromTime))
	buffer.WriteString("\n")

	_, err = f.WriteString(string(buffer.Bytes()))
	if err != nil {
		return err
	}
	f.Close()
	/*
		err := ioutil.WriteFile(filename, buffer.Bytes(), 0644)
	*/

	return nil
}

func floatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}
