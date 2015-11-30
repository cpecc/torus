package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/coreos/pkg/capnslog"
)

func main() {
	app := cli.NewApp()
	capnslog.SetGlobalLogLevel(capnslog.WARNING)
	app.Name = "agroctl"
	app.Usage = "Administer the agro filesystem"
	app.Commands = []cli.Command{
		mkfsCommand,
		listPeersCommand,
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "etcd, C",
			Value: "127.0.0.1:2378",
			Usage: "hostname:port to the etcd instance storing the metadata",
		},
	}
	app.Run(os.Args)
}
