git clone LayerForge in parent folder
Then run to build the wasm

```
GOOS=js GOARCH=wasm go build -o blueprint.wasm main.go
```

Then to host and test run

```
http://localhost:8000/
```