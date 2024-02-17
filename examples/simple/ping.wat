(module
  (import "hosted" "pong" (func $pong (param i32)))
  (; (memory 1) ;)
  (export "ping" (func $ping))
  (global $g i32 (i32.const 0))
  (func $ping
    (global.get $g)
    call $pong
  )
)
