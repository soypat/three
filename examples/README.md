# Examples

## Running examples
Requirements: Install [`wasmserve`](https://github.com/hajimehoshi/wasmserve) and [Go](https://go.dev/). You can install wasmserve with Go:

```shell
go install github.com/hajimehoshi/wasmserve@latest
```

1. First step is to run  inside an example directory (each directory is a self contained project). 
    ```shell
    cd examples/earth
    wasmserve
    ```
2. Head over to where the WASM is being served  [localhost:8080](https://localhost:8080) (default port is :8080 for wasmserve)
