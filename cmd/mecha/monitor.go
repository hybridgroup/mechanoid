package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func monitor(cCtx *cli.Context) error {
	cmd := exec.Command("tinygo", "monitor")
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)
	if err := cmd.Run(); err != nil {
		fmt.Printf("tinygo monitor: %v\n", err)
		os.Exit(1)
	}

	outStr, errStr := stdoutBuf.String(), stderrBuf.String()
	if errStr != "" {
		fmt.Println(errStr)
	} else {
		fmt.Println(outStr)
	}

	return nil
}
