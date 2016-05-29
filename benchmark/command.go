package benchmark

import (
	"fmt"

	"time"

	"github.com/codegangsta/cli"
	"github.com/washraf/runcbm/containers/conctrl"
)

var containerID, bundle, logFile string
var count int

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
	fmt.Printf("Bench Mark Container ID %s \n ", containerID)
	conctrl.StartContainer(containerID, bundle)
	fmt.Println("Sleep for 2 Seconds")
	time.Sleep(time.Second * 2)

	err := Run(containerID, count)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = conctrl.CleanUp(containerID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
