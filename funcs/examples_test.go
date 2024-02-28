package funcs_test

import "github.com/hybridgroup/mechanoid/funcs"

func ExampleF00() {
	_ = funcs.Modules{
		"env": {
			"hello": funcs.F00(func() {
				println("I'm alive!")
			}),
		},
	}
}
