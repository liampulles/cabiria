#!/bin/bash

IN="$1"
BASE="${1##*/}"

rm -rf "/tmp/cabiria/train/$BASE"
mkdir -p "/tmp/cabiria/train/$BASE"
ffmpeg -i "$IN" -r 1 "/tmp/cabiria/train/$BASE/$BASE%06d.png"

echo "/tmp/cabiria/train/$BASE"
