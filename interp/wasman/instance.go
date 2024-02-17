package wasman

import (
	wasmaneng "github.com/c0mm4nd/wasman"
)

type Instance struct {
	instance *wasmaneng.Instance
}

func (i *Instance) Call(name string, args ...interface{}) (interface{}, error) {
	if len(args) == 0 {
		results, _, err := i.instance.CallExportedFunc(name)
		if err != nil {
			return nil, err
		}
		return results, nil
	}

	wargs := make([]uint64, len(args))
	switch args[0].(type) {
	case int32:
		for i, v := range args {
			wargs[i] = uint64(v.(int32))
		}
	case uint32:
		for i, v := range args {
			wargs[i] = uint64(v.(uint32))
		}
	}
	results, _, err := i.instance.CallExportedFunc(name, wargs...)
	if err != nil {
		return nil, err
	}

	return results, nil
}
