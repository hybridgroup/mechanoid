package epsilon

import (
	epsilonlib "github.com/ziggy42/epsilon/epsilon"
)

// Instance implements the engine.Instance interface for Epsilon
type Instance struct {
	instance *epsilonlib.ModuleInstance
}

// Call invokes a function in the Epsilon instance with the given name and arguments
func (i *Instance) Call(name string, args ...any) (any, error) {
	if len(args) == 0 {
		results, err := i.instance.Invoke(name)
		if err != nil {
			return nil, err
		}

		if len(results) == 0 {
			return nil, nil
		}
		if len(results) == 1 {
			return results[0], nil
		}
		return results, nil
	}

	results, err := i.instance.Invoke(name, encodeArgs(args)...)
	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, nil
	}
	if len(results) == 1 {
		return results[0], nil
	}
	return results, nil
}

func encodeArgs(args []any) []any {
	encoded := make([]any, 0, len(args))
	for _, arg := range args {
		encoded = append(encoded, encodeArg(arg))
	}
	return encoded
}

func encodeArg(arg any) any {
	switch val := arg.(type) {
	case int32:
		return val
	case int64:
		return val
	case float32:
		return val
	case float64:
		return val
	case uint32:
		return int32(val)
	case uint64:
		return int64(val)
	case uintptr:
		return int64(val)
	}
	panic("bad arg type")
}
