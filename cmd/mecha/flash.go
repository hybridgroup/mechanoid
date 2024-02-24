package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func flash(cCtx *cli.Context) error {
	if cCtx.Args().Len() < 1 {
		return fmt.Errorf("target board required")
	}

	if cCtx.Bool("monitor") {
		return flashAndMonitor(cCtx)
	}

	targetName := cCtx.Args().First()

	var stdout, stderr bytes.Buffer
	cmd := exec.Command("tinygo", "flash", "-size", "short", "-target", targetName, ".")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("tinygo flash -size short -target %s .: %v\n%s%s", targetName, err, stderr.Bytes(), stdout.Bytes())
	}

	fmt.Println(stdout.String())

	return nil
}

func flashAndMonitor(cCtx *cli.Context) error {
	targetName := cCtx.Args().First()

	var stderr bytes.Buffer
	cmd := exec.Command("tinygo", "flash", "-size", "short", "-target", targetName, "-monitor", ".")
	cmd.Stderr = &stderr
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		log.Fatalf("tinygo flash -size short -target %s .: %v\n%s", targetName, err, stderr.Bytes())
	}

	// print the monitoring output
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}

	return nil
}
