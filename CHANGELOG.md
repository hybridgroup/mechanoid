0.2.0
---

**all**
- update to use main for of wazero with TinyGo changes now merged upstream

**build**
- remove unneeded build setup, and also build on PRs/merge to dev branch

**cmd/mecha**
- add -type flag to mecha create module to support Rust and Zig module templates
- add ability to compile Zig modules to mecha build command
- can build Rust WASM unknown modules
- use main wazero fork for new project creation
- extend build command with project/modules subcommands and build flags

**docs** 
- add info on how to contribute code to the project
- explain that modules can be developed using TinyGo, Rust, or Zig
- make descriptions consistent with site
- now using the main upstream wazero
- remove instructions for GOPRIVATE
- some clarifications on how the Mechanoid architecture works

**interp/wazero** 
- add memory debugging and correct Halt() implementation
- correct handling for Init/Load/Halt to reclaim all resources possible from runtime

0.1.0
---

Initial release
