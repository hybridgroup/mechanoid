package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"

	"github.com/urfave/cli/v2"
)

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
