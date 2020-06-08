set -ex
GOOS=js GOARCH=wasm go build -o static/spaceinvaders.wasm .
cp $(go env GOROOT)/misc/wasm/wasm_exec.js static/wasm_exec.js