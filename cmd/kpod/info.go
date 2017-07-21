package main

import (
	"encoding/json"
	"fmt"

	"github.com/ghodss/yaml"
	"github.com/kubernetes-incubator/cri-o/libkpod/info"
	_ "github.com/kubernetes-incubator/cri-o/libkpod/info/debug"
	_ "github.com/kubernetes-incubator/cri-o/libkpod/info/host"
	_ "github.com/kubernetes-incubator/cri-o/libkpod/info/store"
	"github.com/urfave/cli"
)

var (
	infoDescription = "display system information"
	infoCommand     = cli.Command{
		Name:        "info",
		Usage:       infoDescription,
		Description: `Information display here pertain to the host, current storage stats, and build of kpod. Useful for the user and when reporting issues.`,
		Flags:       infoFlags,
		Action:      infoCmd,
		ArgsUsage:   "",
	}
	infoFlags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, D",
			Usage: "display additional debug information",
		},
		cli.BoolFlag{
			Name:  "json",
			Usage: "output as JSON instead of the default YAML",
		},
	}
)

func infoCmd(c *cli.Context) error {
	i := info.CollectInfo(kc.FromCli(c), info.InfoGivers())

	var buf []byte
	var err error
	if c.Bool("json") {
		buf, err = json.MarshalIndent(i, "", "  ")
	} else {
		buf, err = yaml.Marshal(i)
	}
	if err != nil {
		return err
	}
	fmt.Println(string(buf))

	return nil
}
