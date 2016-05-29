package conutil

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/cpu"
	"github.com/washraf/runcbm/common"
	"github.com/washraf/runcbm/containers/config"
)

//helper https://blog.docker.com/2013/10/gathering-lxc-docker-containers-metrics/

//GetContainerUtilization ...
func GetContainerUtilization(ContainerID string) (ContainerMetrics, error) {
	res := ContainerMetrics{}
	_, ct, err := getCPUUtlizationPercent(ContainerID)
	if err != nil {
		return res, err
	}
	res.CPUTime = ct

	m, err := getUsedMemory(ContainerID)
	if err != nil {
		return res, err
	}
	res.UsedMemory = m

	pcount, err := getProcessesCount(ContainerID)
	if err != nil {
		return res, err
	}
	res.ProcessCount = pcount
	tcount, err := getTaskCount(ContainerID)
	if err != nil {
		return res, err
	}
	res.TaskCount = tcount
	fs, err := getRootFSSize(ContainerID)
	if err != nil {
		return res, err
	}
	res.RootSize = fs

	return res, nil
}

func (m ContainerMetrics) String() string {
	res, _ := json.Marshal(m)
	return (string(res))
}

func getCPUUtlizationPercent(ContainerID string) (float64, uint64, error) {
	c, err := config.ReadControlGroupFile(ContainerID, "cpuacct", "cpuacct.usage")
	if err != nil {
		return 0, 0, err
	}
	conCPU, err := strconv.Atoi(strings.TrimSpace(c))
	if err != nil {
		return 0, 0, err
	}
	fullCPU, err := cpu.Times(false)
	if err != nil {
		return 0, 0, err
	}
	tbusy := fullCPU[0].Total() - fullCPU[0].Idle

	cbusy := float64(conCPU) / 10000000.0
	return common.Round(cbusy/tbusy, 0.5, 3), uint64(cbusy), nil
}

//get memory in Megabytes
func getUsedMemory(containerID string) (int, error) {
	v, err := config.ReadControlGroupFile(containerID, "memory", "memory.usage_in_bytes")
	if err != nil {
		return 0, err
	}
	memory, err := strconv.Atoi(strings.TrimSpace(v))
	if err != nil {
		return 0, err
	}

	return memory / 1048576, nil
}

//get memory in Megabytes
func getProcessesCount(containerID string) (int, error) {
	file, err := config.ReadControlGroupFile(containerID, "cpu", "cgroup.procs")
	if err != nil {
		return 0, err
	}
	arr := strings.Split(file, "\n")
	if err != nil {
		//fmt.Println(err)
		return 0, err
	}

	return len(arr) - 1, nil
}

func getTaskCount(containerID string) (int, error) {
	file, err := config.ReadControlGroupFile(containerID, "cpu", "tasks")
	if err != nil {
		return 0, err
	}
	arr := strings.Split(file, "\n")
	if err != nil {
		//fmt.Println(err)
		return 0, err
	}

	return len(arr) - 1, nil
}

func getRootFSSize(containerID string) (int, error) {
	s, err := config.FindContainerBundle(containerID)
	if err != nil {
		//fmt.Println(err)
		return 0, err
	}
	return common.FindDiskSizeMB(s + "/rootfs")
}
