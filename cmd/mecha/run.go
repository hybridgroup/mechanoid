package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func run(cCtx *cli.Context) error {
	fmt.Println("Running using interpreter", cCtx.String("interpreter"))

	cmd := exec.Command("go", "run", "-tags", cCtx.String("interpreter"), ".")

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

	if err := cmd.Run(); err != nil {
		fmt.Printf("%s: %v\n", cmd.String(), err)
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
