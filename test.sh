#!/usr/bin/env bash

set -e
set -x


########################
#    Golang Testing    #
########################

# https://golang.org/pkg/testing/
go test

# https://golang.org/cmd/vet/
go vet -composites=false

# https://github.com/fzipp/gocyclo
$GOPATH/bin/gocyclo -over 9 .

# https://golang.org/x/lint
$GOPATH/bin/golint .

# https://github.com/gordonklaus/ineffassign
$GOPATH/bin/ineffassign .

# https://github.com/client9/misspell
$GOPATH/bin/misspell -error .


########################
#  Javascript Testing  #
########################

# https://jshint.com/
./node_modules/.bin/jshint --reporter ./.jshintrc.reporter.js ./www/chat/js/

# https://eslint.org/
./node_modules/.bin/eslint --format ./.eslintrc.formatter.js ./www/chat/js/
