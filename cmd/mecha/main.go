package main

import (
	"fmt"
	"os"

	"github.com/hybridgroup/mechanoid"
	"github.com/urfave/cli/v2"
)

var templateFlags = []cli.Flag{
	&cli.StringFlag{Name: "template", Aliases: []string{"t"}, Usage: "template to use for module creation"},
}

func main() {
	app := &cli.App{
		Name:    "mecha",
		Usage:   "Mechanoid CLI",
		Version: mechanoid.Version(),
		Commands: []*cli.Command{
			{
				Name:      "new",
				Usage:     "Create new Mechanoid project or module",
				ArgsUsage: "<name e.g. 'domain.com/projectname'>",
				Action:    createProject,
				Flags:     templateFlags,
				Subcommands: []*cli.Command{
					{
						Name:      "project",
						Usage:     "Create new Mechanoid project",
						ArgsUsage: "<name e.g. 'domain.com/projectname'>",
						Flags:     templateFlags,
						Action:    createProject,
					},
					{
						Name:      "module",
						Usage:     "Create new Mechanoid module",
						ArgsUsage: "<name e.g. 'domain.com/modulename'>",
						Flags:     templateFlags,
						Action:    createModule,
					},
				},
			},
			{
				Name:   "build",
				Usage:  "Build binary files for Mechanoid project",
				Action: build,
			},
			{
				Name:      "flash",
				Usage:     "Flash Mechanoid project to hardware",
				Action:    flash,
				ArgsUsage: "<board-name>",
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "monitor", Aliases: []string{"m"}, Usage: "monitor the serial port after flashing"},
					&cli.StringFlag{Name: "interpreter", Aliases: []string{"i"}, Value: "wazero", Usage: "WebAssembly interpreter to use (wasman, wazero)"},
					&cli.BoolFlag{Name: "debug", Aliases: []string{"d"}, Usage: "perform additional logging for debugging"},
				},
			},
			{
				Name:   "run",
				Usage:  "Run code for Mechanoid project",
				Action: run,
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "interpreter", Aliases: []string{"i"}, Value: "wasman", Usage: "WebAssembly interpreter to use (wasman, wazero)"},
				},
			},
			{
				Name:   "test",
				Usage:  "Run tests for Mechanoid project",
				Action: test,
			},
			{
				Name:   "monitor",
				Usage:  "Monitor connection to hardware using the serial port",
				Action: monitor,
			},
			{
				Name:   "about",
				Usage:  "About Mechanoid",
				Action: about,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
