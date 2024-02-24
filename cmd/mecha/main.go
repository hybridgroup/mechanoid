package main

import (
	"log"
	"os"

	"github.com/hybridgroup/mechanoid"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "mecha",
		Usage:   "Mechanoid CLI",
		Version: mechanoid.Version(),
		Commands: []*cli.Command{
			{
				Name:      "new",
				Usage:     "create new Mechanoid project",
				Args:      true,
				ArgsUsage: "[project-name] <template>",
				Action:    create,
			},
			{
				Name:   "build",
				Usage:  "build Mechanoid project to binary file",
				Action: build,
			},
			{
				Name:      "flash",
				Usage:     "flash Mechanoid project to board",
				Action:    flash,
				ArgsUsage: "[board-name]",
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "monitor", Aliases: []string{"m"}, Usage: "monitor the serial port after flashing"},
				},
			},
			{
				Name:   "test",
				Usage:  "run tests for Mechanoid project",
				Action: test,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
