.PHONY: run
run: lint build
	wasmedge ./main.wasm

.PHONY: lint
lint:
	gofmt -w ./main.go

.PHONY: build
build: clean
	GOOS=wasip1 GOARCH=wasm go build -o main.wasm ./main.go
	wasmedge compile ./main.wasm ./main.wasm

.PHONY: clean
clean:
	rm -f ./main.wasm
