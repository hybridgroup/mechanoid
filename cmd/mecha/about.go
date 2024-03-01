package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var logo = `

  __  __           _                       _     _ 
 |  \/  | ___  ___| |__   __ _ _ __   ___ (_) __| |
 | |\/| |/ _ \/ __| '_ \ / _' | '_ \ / _ \| |/ _' |
 | |  | |  __/ (__| | | | (_| | | | | (_) | | (_| |
 |_|  |_|\___|\___|_| |_|\__,_|_| |_|\___/|_|\__,_|
                                                   

 Mechanoid - Framework for WebAssembly on Embedded Devices 

 https://mechanoid.io
`

func about(cCtx *cli.Context) error {
	fmt.Println(logo)

	return nil
}
