package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"

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

	fmt.Println("Flashing", targetName)

	var cmd *exec.Cmd
	if cCtx.Bool("monitor") {
		cmd = exec.Command("tinygo", "flash", "-size", "short", "-stack-size", "8kb", "-target", targetName, "-monitor", ".")
	} else {
		cmd = exec.Command("tinygo", "flash", "-size", "short", "-stack-size", "8kb", "-target", targetName, ".")
	}

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

	if err := cmd.Run(); err != nil {
		fmt.Printf("tinygo flash -size short -stack-size 8kb -target %s .: %v\n", targetName, err)
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
