module github.com/hybridgroup/mechanoid

go 1.22.0

replace github.com/tetratelabs/wazero => github.com/orsinium-forks/wazero v0.0.0-20240217173836-b12c024bcbe4

require (
	github.com/hybridgroup/wasman v0.0.0-20240304140329-ce1ea6b61834
	github.com/orsinium-labs/wypes v0.1.1
	github.com/tetratelabs/wazero v1.6.0
	github.com/urfave/cli/v2 v2.27.1
	tinygo.org/x/tinyfs v0.3.1-0.20231212053859-32ae3f6bbad9
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
)
