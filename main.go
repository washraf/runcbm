package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/washraf/runcbm/benchmark"
	"github.com/washraf/runcbm/common"
	"github.com/washraf/runcbm/containers/conutil"
)

const usage = `runc benchmarking tool`

func main() {

	defer func() {
		if e := recover(); e != nil {
			if ex, ok := e.(common.Exit); ok == true {
				os.Exit(ex.Code)
			}
			panic(e)
		}
	}()
	app := cli.NewApp()
	app.Name = "runcbm"
	app.Usage = usage
	app.Commands = []cli.Command{
		benchmark.Command,
		conutil.Command,
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
