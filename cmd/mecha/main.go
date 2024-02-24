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
				ArgsUsage: "<name e.g. 'domain.com/projectname'>",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "template", Aliases: []string{"t"}, Usage: "template to use for project creation"},
				},
				Action: createProject,
				Subcommands: []*cli.Command{
					{
						Name:      "project",
						Usage:     "create new Mechanoid project",
						Args:      true,
						ArgsUsage: "<name e.g. 'domain.com/projectname'>",
						Action:    createProject,
					},
					{
						Name:      "module",
						Usage:     "create new Mechanoid module",
						Args:      true,
						ArgsUsage: "<name e.g. 'domain.com/modulename'>",
						Action:    createModule,
					},
				},
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
				ArgsUsage: "<board-name>",
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
