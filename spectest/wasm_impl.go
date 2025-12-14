package spectest

import (
	"github.com/PrismAIO/wasm.go/binary"
	"github.com/PrismAIO/wasm.go/instance"
	"github.com/PrismAIO/wasm.go/interpreter"
	"github.com/PrismAIO/wasm.go/validator"
)

var _ WasmImpl = (*WasmInterpreter)(nil)

type WasmImpl interface {
	Validate(m binary.Module) error
	Instantiate(m binary.Module, instances instance.Map) (instance.Module, error)
	InstantiateBin(data []byte, instances instance.Map) (instance.Module, error)
}

type WasmInterpreter struct{}

func (WasmInterpreter) Validate(m binary.Module) error {
	return validator.Validate(m)
}

func (WasmInterpreter) Instantiate(
	m binary.Module, instances instance.Map,
) (instance.Module, error) {
	return interpreter.New(m, instances)
}

func (WasmInterpreter) InstantiateBin(
	data []byte, instances instance.Map,
) (instance.Module, error) {
	m, err := binary.Decode(data)
	if err != nil {
		return nil, err
	}
	return interpreter.New(m, instances)
}
