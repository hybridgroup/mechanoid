package tester

import (
	_ "embed"
)

//go:embed tester.wasm
var wasmData []byte
