.PHONY: run
run: build
	./server-amd64

.PHONY: runmac
runmac: build
	./server-arm64

.PHONY: build
build: clean wasm_exec
	tinygo build -o ./html/wasm.wasm -target wasm ./wasm.go
	GOOS=linux GOARCH=amd64 go build -o server-amd64 ./server.go
	GOOS=darwin GOARCH=arm64 go build -o server-arm64 ./server.go

.PHONY: wasm_exec
wasm_exec:
	cp `tinygo env TINYGOROOT`/targets/wasm_exec.js ./html/

.PHONY: clean
clean:
	rm -f ./server-amd64
	rm -f ./server-arm64
	rm -f ./html/wasm_exec.js
	rm -f ./html/wasm.wasm
