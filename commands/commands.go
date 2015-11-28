package commands

import (
	"github.com/codegangsta/cli"

	"github.com/yuuki1/dochroot/log"
)

var Commands = []cli.Command{
	CommandPush,
	CommandPull,
	CommandRun,
	CommandUmount,
}

func fatalOnError(command func(context *cli.Context) error) func(context *cli.Context) {
	return func(context *cli.Context) {
		if err := command(context); err != nil {
			log.Error(err)
		}
	}
}
