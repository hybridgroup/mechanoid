module github.com/hybridgroup/mechanoid

go 1.22.0

replace github.com/c0mm4nd/wasman => github.com/hybridgroup/wasman v0.0.0-20240221230704-63fe31eeb0c3

//replace github.com/c0mm4nd/wasman => ../../wasman

require (
	github.com/c0mm4nd/wasman v0.0.0-20220422074058-87e38ef26abd
	tinygo.org/x/tinyfs v0.3.1-0.20231212053859-32ae3f6bbad9
)
