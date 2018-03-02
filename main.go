package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"
)

var eraDict []EraDict

func main() {

	app := cli.NewApp()
	app.Name = "wareki"
	app.Usage = ""
	app.Version = "0.0.1"

	// global options
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "wareki, w",
			Usage: "-w ${YEAR}/${MONTH}/${DAY}",
		},
	}

	// commands
	app.Commands = []cli.Command{
		{
			Name:    "wareki",
			Aliases: []string{"w"},
			Usage:   "wareki -w ${YEAR}/${MONTH}/${DAY}",
			Action:  SeirekiToWarekiAction,
		},
	}

	app.Before = func(c *cli.Context) error {
		var err error
		if eraDict, err = LoadEras(); err != nil {
			return err
		}
		return nil
	}

	app.After = func(c *cli.Context) error {
		return nil
	}

	app.Run(os.Args)
}

func SeirekiToWarekiAction(c *cli.Context) {

	// グローバルオプション
	var toWareki = c.GlobalBool("w")
	if toWareki {
		fmt.Println("this is dry-run")
	}

	var dateStr = ""
	if len(c.Args()) > 0 {
		dateStr = c.Args().First()
	}

	dateSlice := strings.Split(dateStr, "/")
	fmt.Printf("%v\n", dateSlice)
}
