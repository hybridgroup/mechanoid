# Ping

Exports a `ping()` function, that immediately calls the host's `pong()` function.

## Building

```
tinygo build -size short -o ./examples/modules/ping/ping.wasm -target ./examples/modules/ping/wasm-unknown.json -no-debug ./examples/modules/ping
```
