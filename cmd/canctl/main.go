package main

import (
	"os"

	"github.com/ernoaapa/can/cmd"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// Version string to be set at compile time via command line (-ldflags "-X main.VersionString=1.2.3")
var (
	VersionString string
	extraCmds     = []cli.Command{}
)

func main() {
	app := cli.NewApp()
	app.Version = VersionString
	app.Before = cmd.GlobalBefore
	app.Name = "can-cli"
	app.Usage = "commandline tool for managing 'cand'"
	app.Description = `The 'can-cli' is tool for managing 'cand' agent in the device.
	 With this tool, you can create, view and remove containers from the device.`
	app.UsageText = "can-cli [global options] command [command options] [arguments...]"

	app.Flags = append([]cli.Flag{
		cli.StringFlag{
			Name:   "config, c",
			Usage:  "Client configuration",
			EnvVar: "CAN_CONFIG",
			Value:  "~/.can/config",
		},
		cli.StringFlag{
			Name:   "namespace",
			Usage:  "Namespace to use with commands. By default reads from config.",
			EnvVar: "CAN_NAMESPACE",
		},
		cli.StringFlag{
			Name:   "endpoint",
			Usage:  "Use specific device endpoint. E.g. '192.168.1.101:5000'",
			EnvVar: "CAN_ENDPOINT",
		},
	}, cmd.GlobalFlags...)

	app.Commands = []cli.Command{
		getCommand,
		describeCommand,
		deleteCommand,
		attachCommand,
		runCommand,
		createCommand,
		configCommand,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}