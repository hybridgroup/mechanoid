#!/bin/bash

echo $'\nsave ping.wasm 53\n' > /dev/ttyACM0
cat ./ping.wasm > /dev/ttyACM0
