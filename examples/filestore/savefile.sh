#!/bin/bash

# This script is used to save WASM file to a port
# ./savefile.sh <filename> <port>

FILENAME=$1
FILESIZE=$(stat -c%s "$FILENAME")
SHORTNAME=$(basename -- "$FILENAME")
PORT=$2

printf "\r\nsave $SHORTNAME $FILESIZE\r\n" > "$PORT"
cat "$FILENAME" > $PORT
