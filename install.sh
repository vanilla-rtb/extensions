#!/bin/bash

rm -rf $GOPATH/src/github.com/vanilla-rtb/
rm -rf $GOPATH/src/stubs
go get github.com/jessevdk/go-flags
go get -d github.com/vanilla-rtb/extensions
mv $GOPATH/src/github.com/vanilla-rtb/extensions/stubs $GOPATH/src/
go install github.com/vanilla-rtb/extensions
go install stubs

