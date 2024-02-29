package wasman

import (
	"testing"

	"github.com/hybridgroup/mechanoid/interp/tester"
)

func TestName(t *testing.T) {
	i := Interpreter{}
	if i.Name() != "wasman" {
		t.Errorf("Interpreter.Name() failed: %v", i.Name())
	}
}

func TestInit(t *testing.T) {
	tester.InitTest(t, &Interpreter{})
}

func TestLoad(t *testing.T) {
	tester.LoadTest(t, &Interpreter{})
}

func TestRun(t *testing.T) {
	tester.RunTest(t, &Interpreter{})
}

func TestHalt(t *testing.T) {
	tester.HaltTest(t, &Interpreter{})
}

func TestExternalReferences(t *testing.T) {
	tester.ReferencesTest(t, &Interpreter{})
}

func TestDefineFunc(t *testing.T) {
	t.Skip("TODO: implement TestDefineFunc")
}
