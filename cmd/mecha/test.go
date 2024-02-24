package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func test(cCtx *cli.Context) error {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command("go", "test", "./...")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("go test ./...: %v\n%s%s", err, stderr.Bytes(), stdout.Bytes())
	}

	fmt.Println(stdout.String())

	return nil
}
