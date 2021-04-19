#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)

BinaryName="edgex_admin"

echo "$CURDIR/bin/${BinaryName}"

exec $CURDIR/bin/${BinaryName}