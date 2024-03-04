package engine

import "github.com/orsinium-labs/wypes"

type Device interface {
	Init() error
	Modules() wypes.Modules
}
