package benchmark

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/washraf/runcbm/containers/conctrl"
)

var containerID, bundle string
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
	},
}

func benchMark(context *cli.Context) {
	//To Do Test Valid Container ID
	if len(containerID) <= 0 {
		fmt.Println("container id cannot be empty")
		return
	}
	//To Do Test valid container bundle
	if len(bundle) <= 0 {
		fmt.Println("bundle id cannot be empty")

		return
	}
	fmt.Printf("Bench Mark Container ID %s \n ", containerID)
	conctrl.StartContainer(containerID, bundle)
	err := Run(containerID, count)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = conctrl.CleanUp(containerID)
	if err != nil {
		fmt.Println(err)
		return
	}
}
