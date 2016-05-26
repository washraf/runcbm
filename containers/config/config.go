package config

import (
	"io/ioutil"
	"strings"

	"github.com/washraf/runcbm/common"
)

const (
	containersbase = "/run/runc/"
	cgroupbase     = "/sys/fs/cgroup/"
	file           = "/state.json"
)

func readStateFile(ContainerID string) ([]byte, error) {
	add := containersbase + ContainerID + file
	return ioutil.ReadFile(add)
}

//FindContainerBundle ...
func FindContainerBundle(ContainerID string) (string, error) {
	full, err := readStateFile(ContainerID)
	if err != nil {
		return "", err
	}
	//I Don't need containerId Any more
	fullrootfs, err := common.GetItemFromJSON(full, "config", "rootfs")
	if err != nil {
		return "", err
	}
	a := strings.Split(fullrootfs, "/")
	rootfs := strings.Join(a[:len(a)-1], "/")
	return rootfs, nil
}

//FindCGroupPath ...
func FindCGroupPath(ContainerID string) (string, error) {
	full, err := readStateFile(ContainerID)
	if err != nil {
		return "", err
	}
	//I Don't need containerId Any more
	path, err := common.GetItemFromJSON(full, "config", "cgroups", "path")
	if err != nil {
		return "", err
	}
	return path, nil
}

//ReadControlGroupFile ...
func ReadControlGroupFile(containerID, cgroup, item string) (string, error) {
	path, err := FindCGroupPath(containerID)
	if err != nil {
		return "", err
	}
	add := cgroupbase + cgroup + path + "/" + item
	r, err := ioutil.ReadFile(add)
	return string(r), err
}
