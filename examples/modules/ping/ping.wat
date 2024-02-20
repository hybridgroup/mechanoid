(module
  (type (;0;) (func))
  (import "hosted" "pong" (func $main.pong (type 0)))
  (func $__wasm_call_dtors (type 0))
  (func $ping (type 0)
    call $main.pong)
  (export "_initialize" (func $__wasm_call_dtors))
  (export "ping" (func $ping)))
