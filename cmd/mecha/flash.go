package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"

	"github.com/briandowns/spinner"
	"github.com/urfave/cli/v2"
)

func flash(cCtx *cli.Context) error {
	if cCtx.Args().Len() < 1 {
		return fmt.Errorf("target board required")
	}

	// build all the modules before flashing the hardware
	if err := build(cCtx); err != nil {
		return err
	}

	targetName := cCtx.Args().First()

	s := spinner.New(spinner.CharSets[17], 100*time.Millisecond, spinner.WithWriter(os.Stdout))
	s.Suffix = " Building application for " + targetName + " using interpreter " + cCtx.String("interpreter")
	s.FinalMSG = "Application built. Now flashing...\n"
	s.Start()
	defer s.Stop()

	intp := cCtx.String("interpreter")
	if intp == "wasman" {
		intp = "wasman nowazero"
	}

	if cCtx.Bool("debug") {
		intp += " debug"
	}

	var cmd *exec.Cmd
	if cCtx.Bool("monitor") {
		cmd = exec.Command("tinygo", "flash", "-size", "short", "-stack-size", "8kb", "-tags", intp, "-target", targetName, "-monitor", ".")
	} else {
		cmd = exec.Command("tinygo", "flash", "-size", "short", "-stack-size", "8kb", "-tags", intp, "-target", targetName, ".")
	}

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(&spinWriter{s, os.Stdout, false}, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(&spinWriter{s, os.Stderr, false}, &stderrBuf)

	if err := cmd.Run(); err != nil {
		fmt.Printf("%s: %v\n", cmd.String(), err)
		os.Exit(1)
	}

	// print the monitoring output
	if cCtx.Bool("monitor") {
		outStr, errStr := stdoutBuf.String(), stderrBuf.String()
		if errStr != "" {
			fmt.Println(errStr)
		} else {
			fmt.Println(outStr)
		}
	}

	return nil
}
