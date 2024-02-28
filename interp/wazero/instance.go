package wazero

import (
	"context"
	"errors"

	"github.com/tetratelabs/wazero/api"
)

type Instance struct {
	module api.Module
}

func (i *Instance) Call(name string, args ...any) (any, error) {
	f := i.module.ExportedFunction(name)
	if f == nil {
		return nil, errors.New("Function not found")
	}
	ctx := context.Background()
	rawResults, err := f.Call(ctx, encodeArgs(args)...)
	if err != nil {
		return nil, err
	}
	results := decodeResults(rawResults, f.Definition().ResultTypes())
	if len(results) == 0 {
		return nil, nil
	}
	if len(results) == 1 {
		return results[0], nil
	}
	return results, nil
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
		return api.EncodeI32(val)
	case int64:
		return api.EncodeI64(val)
	case float32:
		return api.EncodeF32(val)
	case float64:
		return api.EncodeF64(val)
	case uint32:
		return api.EncodeU32(val)
	case uint64:
		return uint64(val)
	case uintptr:
		return api.EncodeExternref(val)
	}
	panic("bad arg type")
}

func decodeResults(results []uint64, vtypes []api.ValueType) []any {
	decoded := make([]any, 0, len(results))
	for i, result := range results {
		vtype := vtypes[i]
		decoded = append(decoded, decodeResult(result, vtype))
	}
	return decoded
}

func decodeResult(result uint64, vtype api.ValueType) any {
	switch vtype {
	case api.ValueTypeF32:
		return api.DecodeF32(result)
	case api.ValueTypeF64:
		return api.DecodeF64(result)
	case api.ValueTypeI32:
		return api.DecodeI32(result)
	case api.ValueTypeI64:
		return int64(result)
	case api.ValueTypeExternref:
		return api.DecodeExternref(result)
	}
	panic("unreachable")
}
