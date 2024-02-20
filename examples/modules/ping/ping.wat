(module
  (import "hosted" "pong" (func $pong))
  (; (memory 1) ;)
  (export "ping" (func $ping))
  (func $ping
    call $pong
  )
)
