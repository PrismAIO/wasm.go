
wastgo:
	@[[ -d spec ]] || git clone --depth=10 --branch=master https://github.com/WebAssembly/spec
	@go build github.com/PrismAIO/wasm.go/cmd/wasmgo

test: wasmgo
	@sh ./run_testsuite.sh

.PHONY: clean

clean:
	@rm -rf wasmgo
