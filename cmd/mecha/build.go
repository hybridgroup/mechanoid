package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func build(cCtx *cli.Context) error {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	if stat, err := os.Stat(filepath.Join(wd, "modules")); err == nil && stat.IsDir() {

		// build modules?
		dirs, err := os.ReadDir(filepath.Join(wd, "modules"))
		if err != nil {
			log.Fatal(err)
		}

		for _, dir := range dirs {
			f, err := dir.Info()
			if err != nil {
				log.Fatal(err)
			}
			if !f.IsDir() {
				continue
			}
			if err := buildModule(filepath.Join(wd, "modules"), f.Name()); err != nil {
				return err
			}
		}
	}

	// TODO: build main project

	return nil
}

func buildModule(modulesPath, name string) error {
	modulePath := filepath.Join(modulesPath, name)
	os.Chdir(modulePath)
	defer os.Chdir("../..")

	var stdout, stderr bytes.Buffer
	cmd := exec.Command("tinygo", "build", "-o", filepath.Join(modulesPath, name+".wasm"), "-target", name+".json", "-no-debug", ".")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("tinygo build error %s: %v\n%s%s", modulePath, err, stderr.Bytes(), stdout.Bytes())
	}

	return nil
}
