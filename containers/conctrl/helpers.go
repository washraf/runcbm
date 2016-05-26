package conctrl

import (
	"fmt"
	"os/exec"
	"time"
)

//StartContainer Starts a Container at the start of the simmulation
func StartContainer(containerID, dir string) error {
	fmt.Println("Starting Container")
	command := exec.Command("time", "-f", "%e", "runc", "start", "-d", containerID)
	command.Dir = dir
	err := command.Start()
	if err != nil {
		return err
	}
	time.Sleep(time.Second * 1)
	return nil
}

//CleanUp Kills and deletes the container
func CleanUp(containerID string) error {
	if e := kill(containerID); e != nil {
		return e
	}
	if e := delete(containerID); e != nil {
		return e
	}
	return nil
}

func kill(containerID string) error {
	fmt.Println("Kill Container")
	command := exec.Command("runc", "kill", containerID)
	err := command.Start()
	if err != nil {
		return err
	}
	time.Sleep(time.Second * 1)
	return nil
}

func delete(containerID string) error {
	fmt.Println("Delete Container")
	command := exec.Command("runc", "delete", containerID)
	err := command.Start()
	if err != nil {
		return err
	}
	time.Sleep(time.Second * 1)
	return nil
}

//TestRunning Check if the Container is running or not
func TestRunning(ContainerID string) {
	fmt.Println("TestRunning")
}
