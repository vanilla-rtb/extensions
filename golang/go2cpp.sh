#!/bin/bash

export CGO_LDFLAGS="-Wl,--start-group "$(< /dev/stdin)" -Wl,--end-group"

echo CGO_LDFLAGS=${CGO_LDFLAGS}


echo CGO_LDFLAGS_ALLOW="-Wl,-unresolved-symbols=ignore-all" go build  -buildmode=exe -o $1 $2

CGO_LDFLAGS_ALLOW="-Wl,-unresolved-symbols=ignore-all" go build  -buildmode=exe -o $1 $2

