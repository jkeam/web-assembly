.PHONY: run
run: build
	wasmedge ./main.wasm

.PHONY: build
build: clean
	tinygo build -o ./main.wasm -target wasi ./main.go
	wasmedge compile ./main.wasm ./main.wasm

.PHONY: clean
clean:
	rm -f ./main.wasm
