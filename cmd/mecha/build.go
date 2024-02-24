package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func build(cCtx *cli.Context) error {
	fmt.Println("build: ", cCtx.Args().First())
	return nil
}
