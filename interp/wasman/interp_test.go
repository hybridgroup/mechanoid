package wasman

import (
	"testing"
)

func TestName(t *testing.T) {
	i := Interpreter{}
	if i.Name() != "wasman" {
		t.Errorf("Interpreter.Name() failed: %v", i.Name())
	}
}

func TestInit(t *testing.T) {
	i := Interpreter{}
	err := i.Init()
	if err != nil {
		t.Errorf("Interpreter.Init() failed: %v", err)
	}
}

func TestLoad(t *testing.T) {
	t.Skip("TODO: implement TestLoad")
}

func TestRun(t *testing.T) {
	t.Skip("TODO: implement TestRun")
}

func TestHalt(t *testing.T) {
	t.Skip("TODO: implement TestHalt")
}

func TestDefineFunc(t *testing.T) {
	t.Skip("TODO: implement TestDefineFunc")
}
