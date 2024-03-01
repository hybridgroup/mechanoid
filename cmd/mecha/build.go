package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func build(cCtx *cli.Context) error {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if stat, err := os.Stat(filepath.Join(wd, "modules")); err == nil && stat.IsDir() {

		// build modules?
		dirs, err := os.ReadDir(filepath.Join(wd, "modules"))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, dir := range dirs {
			f, err := dir.Info()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
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
	fmt.Println("Building module", name)
	modulePath := filepath.Join(modulesPath, name)
	os.Chdir(modulePath)
	defer os.Chdir("../..")

	cmd := exec.Command("tinygo", "build", "-o", filepath.Join(modulesPath, name+".wasm"), "-target", name+".json", "-no-debug", "-size", "short", ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("tinygo build error %s: %v\n", modulePath, err)
		os.Exit(1)
	}

	return nil
}
