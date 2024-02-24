package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

const (
	defaultProjectTemplate = "github.com/hybridgroup/mechanoid-examples/simple"
	defaultModuleTemplate  = "github.com/hybridgroup/mechanoid-examples/modules/hello"
)

// create a new project or module
func create(cCtx *cli.Context) error {
	switch {
	case cCtx.Args().Len() == 0:
		return fmt.Errorf("new project name required")
	case cCtx.Args().Len() == 1:
		name := cCtx.Args().Get(0)
		templateName := defaultProjectTemplate

		return createFromTemplate(templateName, name)
	case cCtx.Args().Len() == 2:
		switch cCtx.Args().Get(0) {
		case "project":
			name := cCtx.Args().Get(1)
			templateName := defaultProjectTemplate

			return createFromTemplate(templateName, name)
		case "module":
			name := cCtx.Args().Get(1)
			templateName := defaultModuleTemplate

			err := os.MkdirAll("modules", 0777)
			if err != nil {
				log.Fatal(err)
			}
			os.Chdir("modules")
			defer os.Chdir("..")

			return createFromTemplate(templateName, name)
		default:
			return fmt.Errorf("don't know how to create a new: %s", cCtx.Args().Get(0))
		}
	case cCtx.Args().Len() == 3:
		switch cCtx.Args().Get(0) {
		case "project":
			name := cCtx.Args().Get(1)
			templateName := cCtx.Args().Get(2)

			return createFromTemplate(templateName, name)
		case "module":
			name := cCtx.Args().Get(1)
			templateName := cCtx.Args().Get(2)

			err := os.MkdirAll("modules", 0777)
			if err != nil {
				log.Fatal(err)
			}
			os.Chdir("modules")
			defer os.Chdir("..")

			return createFromTemplate(templateName, name)
		default:
			return fmt.Errorf("don't know how to create a new: %s", cCtx.Args().Get(0))
		}
	default:
		return fmt.Errorf("don't know how to create a new: %v", cCtx.Args().Slice())
	}
}

func createFromTemplate(templ, proj string) error {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command("gonew", templ, proj)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("gonew %s %s: %v\n%s%s", templ, proj, err, stderr.Bytes(), stdout.Bytes())
	}

	return nil
}
