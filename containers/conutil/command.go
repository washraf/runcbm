package conutil

import (
	"fmt"

	"github.com/codegangsta/cli"
)

var containerID string

//Command the command it self
var Command = cli.Command{
	Name:      "measure",
	Usage:     "measure container load",
	ArgsUsage: "COMMAND [arguments...]",
	Action:    measureContainer,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "id",
			Usage:       "Container ID",
			Destination: &containerID,
		},
	},
}

func measureContainer(context *cli.Context) error {
	if len(containerID) <= 0 {
		fmt.Println("container id cannot be empty")
		return nil
	}
	fmt.Printf("Measure Container ID %s \n ", containerID)

	u, err := GetContainerUtilization(containerID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(u)
	return nil
}
