#!/bin/bash

go install github.com/hunchly/funchly/onionscan@latest
cd gopath
cp -r `find ~+ -type d -name "templates"` $HOME
cd -
