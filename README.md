# WebAssembly

## WASM

### Prerequisite

1. TinyGo v0.35.0+
2. Go v1.21.11
3. Podman v5+

### Running

```shell
cd ./wasm
make run
# on apple silicon run below
# make runmac
```

## WASI

### Prerequisite

1. WasmEdge v0.14.1+
2. TinyGo v0.35.0+
3. Go v1.21.11
4. Podman v5+

### Running

```shell
cd ./wasi
ls
```

For comparison of the WASI implementation versus regular Go compare these two:

1. `main.go` - WASI implementation
2. `server.go` - Regular Go implementation

For a regular Go application, run:

```shell
go run ./server.go
```

To build and run a WASI app, run:

```shell
make
```

### Testing

After the server is started using either from above,
open a new terminal and run:

```shell
curl -X POST -d '{"operand1": 1, "operand2": 2}' http://localhost:8080/add
```

### WASI References

1. [WASI and Go](https://go.dev/blog/wasi)
2. [WASI and Go not working without special code](https://github.com/dispatchrun/net/issues/37)
3. [WASI and Go example code](https://github.com/ldemailly/go-scratch/blob/main/tinyhttp/tinyhttp.go)
4. [WasmEdge and CRUN](https://wasmedge.org/docs/develop/deploy/intro#with-crun)
