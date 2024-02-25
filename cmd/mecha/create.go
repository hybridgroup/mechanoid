package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v2"
)

const (
	defaultProjectTemplate = "github.com/hybridgroup/mechanoid-templates/projects/simple"
	defaultModuleTemplate  = "github.com/hybridgroup/mechanoid-templates/modules/hello"
)

func createProject(cCtx *cli.Context) error {
	if cCtx.Args().Len() < 1 {
		return fmt.Errorf("name required")
	}
	name := cCtx.Args().Get(0)
	templateName := cCtx.String("template")
	switch {
	case templateName == "":
		templateName = defaultProjectTemplate
	case !strings.Contains(templateName, "/"):
		templateName = "github.com/hybridgroup/mechanoid-templates/projects/" + templateName
	}

	return createFromTemplate(templateName, name)
}

func createModule(cCtx *cli.Context) error {
	if cCtx.Args().Len() < 1 {
		return fmt.Errorf("name required")
	}
	name := cCtx.Args().Get(0)
	basename := filepath.Base(name)
	templateName := cCtx.String("template")
	switch {
	case templateName == "":
		templateName = defaultModuleTemplate
	case !strings.Contains(templateName, "/"):
		templateName = "github.com/hybridgroup/mechanoid-templates/modules/" + templateName
	}

	err := os.MkdirAll("modules", 0777)
	if err != nil {
		log.Fatal(err)
	}
	os.Chdir("modules")
	defer os.Chdir("..")

	if err := createFromTemplate(templateName, name); err != nil {
		return err
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return os.Rename(filepath.Join(wd, basename, filepath.Base(templateName)+".json"),
		filepath.Join(wd, basename, basename+".json"))
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