package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/briandowns/spinner"
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
	s := spinner.New(spinner.CharSets[17], 100*time.Millisecond, spinner.WithWriter(os.Stdout))
	s.Suffix = " Building module " + name
	s.FinalMSG = "Done.\n"
	s.Start()
	defer s.Stop()

	fmt.Println("Building module", name)
	modulePath := filepath.Join(modulesPath, name)
	os.Chdir(modulePath)
	defer os.Chdir("../..")

	cmd := exec.Command("tinygo", "build", "-o", filepath.Join(modulesPath, name+".wasm"), "-target", name+".json", "-no-debug", "-size", "short", ".")

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(&spinWriter{s, os.Stdout, false}, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(&spinWriter{s, os.Stderr, false}, &stderrBuf)

	if err := cmd.Run(); err != nil {
		fmt.Printf("tinygo build error %s: %v\n", modulePath, err)
		os.Exit(1)
	}

	return nil
}
