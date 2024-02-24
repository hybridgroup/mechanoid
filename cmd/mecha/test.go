package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func testProject(cCtx *cli.Context) error {
	fmt.Println("test: ", cCtx.Args().First())
	return nil
}
