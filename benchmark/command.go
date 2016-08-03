package benchmark

import (
	"fmt"
	"os/exec"
	"strconv"

	"time"

	"github.com/codegangsta/cli"
	"github.com/washraf/runcbm/containers/conctrl"
)

var containerID, bundle, logFile, other string
var count, move int

//Command the command it self
var Command = cli.Command{
	Name:      "bm",
	Usage:     "benchmark a container ",
	ArgsUsage: "COMMAND [arguments...]",
	Action:    benchMark,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "id",
			Usage:       "Container ID",
			Destination: &containerID,
		},
		cli.IntFlag{
			Name:        "n",
			Usage:       "Number of Trails",
			Destination: &count,
			Value:       5,
		},
		cli.StringFlag{
			Name:        "dir",
			Usage:       "location of the bundle folder",
			Destination: &bundle,
		},
		cli.StringFlag{
			Name:        "log",
			Usage:       "location of the log file",
			Destination: &logFile,
			Value:       "/containers/log",
		},
		cli.IntFlag{
			Name:        "move",
			Usage:       "0 for no move; 3 for trial 3 & 4 for trial 4",
			Value:       0,
			Destination: &move,
		},
		cli.StringFlag{
			Name:        "other",
			Usage:       "location for move and back",
			Destination: &other,
			Value:       "/coniscsi/",
		},
	},
}

func benchMark(context *cli.Context) error {
	//To Do Test Valid Container ID
	if len(containerID) <= 0 {
		fmt.Println("container id cannot be empty")
		return nil
	}
	//To Do Test valid container bundle
	if len(bundle) <= 0 {
		fmt.Println("bundle id cannot be empty")

		return nil
	}
	for d := 1; d <= 40; d++ {

		fmt.Println("Create config for iteration " + strconv.Itoa(d))
		err := setProcessesMemory(d, bundle)
		if err != nil {
			fmt.Println("cannot create config")
			return err
		}
		fmt.Printf("Bench Mark Container ID %s \n ", containerID)

		conctrl.StartContainer(containerID, bundle)
		fmt.Println("Sleep for 2 Seconds")
		time.Sleep(time.Second * 2)

		err = Run(d, containerID, count, move, other)
		if err != nil {
			fmt.Println(err)
			return err
		}
		err = conctrl.CleanUp(containerID)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

func setProcessesMemory(i int, condir string) error {
	command := exec.Command("ocitools", "generate",
		"--args", "stress",
		"--args", "-c", "--args", "8",
		"--args", "-m", "--args", strconv.Itoa(i*2),
		"--args", "--vm-bytes", "--args", "25M",
		"--args", "-t", "--args", "60s")
	command.Dir = condir
	_, err := command.CombinedOutput()
	return err
}
