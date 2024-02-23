package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"

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
				Usage:     "create a new Mechanoid project",
				Args:      true,
				ArgsUsage: "<project-name> <template>",
				Action:    newProject,
			},
			{
				Name:   "build",
				Usage:  "build a Mechanoid project to a binary file",
				Action: buildProject,
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

const defaultTemplate = "github.com/hybridgroup/mechanoid-examples/simple"

func newProject(cCtx *cli.Context) error {
	if cCtx.Args().Len() < 1 {
		return fmt.Errorf("new project name required")
	}

	projectName := cCtx.Args().First()
	templateName := defaultTemplate
	if cCtx.Args().Len() > 1 {
		templateName = cCtx.Args().Get(1)
	}

	var stdout, stderr bytes.Buffer
	cmd := exec.Command("gonew", templateName, projectName)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("gonew %s %s: %v\n%s%s", templateName, projectName, err, stderr.Bytes(), stdout.Bytes())
	}

	return nil
}

func buildProject(cCtx *cli.Context) error {
	fmt.Println("build: ", cCtx.Args().First())
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
