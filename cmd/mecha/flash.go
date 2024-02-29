package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func flash(cCtx *cli.Context) error {
	if cCtx.Args().Len() < 1 {
		return fmt.Errorf("target board required")
	}

	targetName := cCtx.Args().First()

	var cmd *exec.Cmd
	if cCtx.Bool("monitor") {
		cmd = exec.Command("tinygo", "flash", "-size", "short", "-target", targetName, "-monitor", ".")
	} else {
		cmd = exec.Command("tinygo", "flash", "-size", "short", "-target", targetName, ".")
	}

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

	if err := cmd.Run(); err != nil {
		log.Fatalf("tinygo flash -size short -target %s .: %v\n", targetName, err) //, stderr.Bytes())
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
