package wasman

import (
	"math"

	wasmaneng "github.com/hybridgroup/wasman"
	"github.com/hybridgroup/wasman/types"
)

type Instance struct {
	instance *wasmaneng.Instance
}

func (i *Instance) Call(name string, args ...any) (any, error) {
	if len(args) == 0 {
		results, types, err := i.instance.CallExportedFunc(name)
		if err != nil {
			return nil, err
		}
		res := decodeResults(results, types)
		if len(res) == 0 {
			return nil, nil
		}
		if len(res) == 1 {
			return res[0], nil
		}
		return res, nil
	}

	results, types, err := i.instance.CallExportedFunc(name, encodeArgs(args)...)
	if err != nil {
		return nil, err
	}

	res := decodeResults(results, types)
	if len(res) == 0 {
		return nil, nil
	}
	if len(res) == 1 {
		return res[0], nil
	}
	return res, nil
}

func encodeArgs(args []any) []uint64 {
	encoded := make([]uint64, 0, len(args))
	for _, arg := range args {
		encoded = append(encoded, encodeArg(arg))
	}
	return encoded
}

func encodeArg(arg any) uint64 {
	switch val := arg.(type) {
	case int32:
		return uint64(val)
	case int64:
		return uint64(val)
	case float32:
		return uint64(math.Float32bits(val))
	case float64:
		return uint64(math.Float64bits(val))
	case uint32:
		return uint64(val)
	case uint64:
		return uint64(val)
	}
	panic("bad arg type")
}

func decodeResults(results []uint64, vtypes []types.ValueType) []any {
	decoded := make([]any, 0, len(results))
	for i, result := range results {
		vtype := vtypes[i]
		decoded = append(decoded, decodeResult(result, vtype))
	}
	return decoded
}

func decodeResult(result uint64, vtype types.ValueType) any {
	switch vtype {
	case types.ValueTypeF32:
		return math.Float32frombits(uint32(result))
	case types.ValueTypeF64:
		return math.Float64frombits(uint64(result))
	case types.ValueTypeI32:
		return int32(result)
	case types.ValueTypeI64:
		return int64(result)
	}
	panic("unreachable")
}
