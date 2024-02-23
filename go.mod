module github.com/hybridgroup/mechanoid

go 1.22.0

replace github.com/c0mm4nd/wasman => github.com/hybridgroup/wasman v0.0.0-20240221230704-63fe31eeb0c3

replace github.com/tetratelabs/wazero => github.com/orsinium-forks/wazero v0.0.0-20240217173836-b12c024bcbe4

//replace github.com/c0mm4nd/wasman => ../../wasman

require (
	github.com/c0mm4nd/wasman v0.0.0-20220422074058-87e38ef26abd
	github.com/urfave/cli/v2 v2.27.1
	tinygo.org/x/tinyfs v0.3.1-0.20231212053859-32ae3f6bbad9
	github.com/tetratelabs/wazero v1.6.0
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
)
