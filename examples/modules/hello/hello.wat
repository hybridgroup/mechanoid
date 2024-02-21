(module
  (type (;0;) (func (param i32 i32) (result i32)))
  (type (;1;) (func))
  (type (;2;) (func (param i32) (result i32)))
  (type (;3;) (func (param i32 i32 i32 i32 i32 i32)))
  (import "hosted" "hola" (func $main.hola (type 0)))
  (import "env" "memory" (memory (;0;) 1 1))
  (func $__wasm_init_memory (type 1)
    i32.const 4136
    i32.const 0
    i32.const 36
    memory.fill)
  (func $_initialize (type 1)
    i32.const 4136
    memory.size
    i32.const 16
    i32.shl
    i32.store
    i32.const 4136
    memory.size
    i32.const 16
    i32.shl
    i32.store)
  (func $runtime.alloc (type 2) (param i32) (result i32)
    (local i32 i32 i32)
    i32.const 4152
    i32.const 4152
    i64.load
    i64.const 1
    i64.add
    i64.store
    i32.const 4128
    i32.const 4128
    i32.load
    local.tee 1
    local.get 0
    i32.const 15
    i32.add
    i32.const -16
    i32.and
    local.tee 2
    i32.add
    local.tee 0
    i32.store
    i32.const 4144
    i32.const 4144
    i64.load
    local.get 2
    i64.extend_i32_u
    i64.add
    i64.store
    i32.const 4136
    i32.load
    local.set 3
    block  ;; label = @1
      loop  ;; label = @2
        local.get 0
        local.get 3
        i32.lt_u
        br_if 1 (;@1;)
        memory.size
        memory.grow
        i32.const -1
        i32.ne
        if  ;; label = @3
          i32.const 4136
          memory.size
          i32.const 16
          i32.shl
          local.tee 3
          i32.store
          i32.const 4128
          i32.load
          local.set 0
          br 1 (;@2;)
        end
      end
      unreachable
    end
    local.get 1
    i32.const 0
    local.get 2
    memory.fill
    local.get 1)
  (func $runtime.sliceAppend (type 3) (param i32 i32 i32 i32 i32 i32)
    (local i32 i32)
    block  ;; label = @1
      local.get 5
      i32.eqz
      if  ;; label = @2
        local.get 1
        local.set 6
        local.get 3
        local.set 7
        br 1 (;@1;)
      end
      block  ;; label = @2
        local.get 4
        local.get 3
        local.get 5
        i32.add
        local.tee 7
        i32.ge_u
        if  ;; label = @3
          local.get 1
          local.set 6
          br 1 (;@2;)
        end
        i32.const 1
        local.get 4
        i32.const 1
        i32.shl
        local.tee 4
        local.get 4
        i32.const 1
        i32.le_u
        select
        local.set 6
        loop  ;; label = @3
          local.get 6
          local.tee 4
          i32.const 1
          i32.shl
          local.set 6
          local.get 4
          local.get 7
          i32.lt_u
          br_if 0 (;@3;)
        end
        local.get 4
        call $runtime.alloc
        local.set 6
        local.get 3
        i32.eqz
        br_if 0 (;@2;)
        local.get 6
        local.get 1
        local.get 3
        memory.copy
      end
      local.get 3
      local.get 6
      i32.add
      local.get 2
      local.get 5
      memory.copy
    end
    local.get 0
    local.get 4
    i32.store offset=8
    local.get 0
    local.get 7
    i32.store offset=4
    local.get 0
    local.get 6
    i32.store)
  (func $hello (type 0) (param i32 i32) (result i32)
    (local i32 i32)
    global.get 0
    i32.const 32
    i32.sub
    local.tee 2
    global.set 0
    i32.const 4160
    i32.const 0
    call $runtime.alloc
    local.tee 3
    i32.store
    i32.const 4168
    i32.const 0
    i32.store
    block  ;; label = @1
      local.get 0
      i32.load offset=8
      local.get 1
      i32.lt_u
      br_if 0 (;@1;)
      local.get 2
      i32.const 16
      i32.add
      local.get 3
      local.get 0
      i32.load
      i32.const 0
      i32.const 0
      local.get 1
      call $runtime.sliceAppend
      i32.const 4160
      local.get 2
      i32.load offset=16
      local.tee 0
      i32.store
      i32.const 4168
      local.get 2
      i32.load offset=20
      local.tee 1
      i32.store
      local.get 2
      local.get 0
      i32.const 4096
      local.get 1
      local.get 2
      i32.load offset=24
      i32.const 30
      call $runtime.sliceAppend
      i32.const 4160
      local.get 2
      i32.load
      local.tee 1
      i32.store
      i32.const 4168
      local.get 2
      i32.load offset=4
      local.tee 0
      i32.store
      local.get 0
      i32.eqz
      br_if 0 (;@1;)
      local.get 1
      local.get 0
      call $main.hola
      drop
      i32.const 4168
      i32.load
      local.get 2
      i32.const 32
      i32.add
      global.set 0
      return
    end
    unreachable)
  (global (;0;) (mut i32) (i32.const 4096))
  (export "_initialize" (func $_initialize))
  (export "hello" (func $hello))
  (start $__wasm_init_memory)
  (data (;0;) (i32.const 4096) " back from TinyGo WebAssembly!")
  (data (;1;) (i32.const 4128) "P\10\00\00"))
