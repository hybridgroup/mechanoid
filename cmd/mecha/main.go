package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "mecha",
		Usage: "Mechanoid WASM embedded development tools",
		Commands: []*cli.Command{
			{
				Name:   "new",
				Usage:  "create a new Mechanoid project",
				Action: newProject,
			},
			{
				Name:   "flash",
				Usage:  "flash a Mechanoid project to a device",
				Action: flashProject,
			},
			{
				Name:   "test",
				Usage:  "run tests for a Mechanoid project",
				Action: testProject,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func newProject(cCtx *cli.Context) error {
	fmt.Println("new: ", cCtx.Args().First())
	return nil
}

func flashProject(cCtx *cli.Context) error {
	fmt.Println("flash: ", cCtx.Args().First())
	return nil
}

func testProject(cCtx *cli.Context) error {
	fmt.Println("test: ", cCtx.Args().First())
	return nil
}
