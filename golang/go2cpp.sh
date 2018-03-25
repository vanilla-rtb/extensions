#!/bin/bash

export CGO_LDFLAGS=$(< /dev/stdin)

echo CGO_LDFLAGS=${CGO_LDFLAGS}

go build  -buildmode=exe -o $1 $2

