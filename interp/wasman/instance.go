package wasman

import (
	wasmaneng "github.com/c0mm4nd/wasman"
)

type Instance struct {
	instance *wasmaneng.Instance
}

func (i *Instance) Call(name string, args ...interface{}) error {
	_, _, err := i.instance.CallExportedFunc(name)
	if err != nil {
		return err
	}

	return nil
}
