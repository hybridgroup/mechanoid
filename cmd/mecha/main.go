package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	templateFlags = []cli.Flag{
		&cli.StringFlag{Name: "template", Aliases: []string{"t"}, Usage: "template to use for module creation"},
		&cli.StringFlag{Name: "type", Value: "tinygo", Usage: "template type [tinygo, rust, zig]"},
	}

	buildFlags = []cli.Flag{
		&cli.StringFlag{Name: "interpreter", Aliases: []string{"i"}, Value: "wazero", Usage: "WebAssembly interpreter to use (wasman, wazero)"},
		&cli.BoolFlag{Name: "debug", Aliases: []string{"d"}, Usage: "perform additional logging for debugging"},
		&cli.StringSliceFlag{
			Name:  "params",
			Usage: "Pass build-time parameters for the application or modules. Format: -params main.name=value -params main.descript=value2",
		},
	}
)

func main() {
	app := &cli.App{
		Name:    "mecha",
		Usage:   "Mechanoid CLI",
		Version: Version(),
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
				Name:      "build",
				Usage:     "Build binary files for Mechanoid project and/or modules",
				ArgsUsage: "<project|modules>",
				Flags:     buildFlags,
				Action:    buildModules,
				Subcommands: []*cli.Command{
					{
						Name:   "project",
						Usage:  "Build current Mechanoid project",
						Action: buildProject,
						Flags:  buildFlags,
					},
					{
						Name:   "modules",
						Usage:  "Build current Mechanoid modules",
						Action: buildModules,
						Flags:  buildFlags,
					},
				},
			},
			{
				Name:      "flash",
				Usage:     "Flash Mechanoid project to hardware",
				Action:    flash,
				ArgsUsage: "<board-name>",
				Flags: append(buildFlags,
					&cli.BoolFlag{Name: "monitor", Aliases: []string{"m"}, Usage: "monitor the serial port after flashing"},
				),
			},
			{
				Name:   "run",
				Usage:  "Run code for Mechanoid project",
				Action: run,
				Flags:  buildFlags,
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
