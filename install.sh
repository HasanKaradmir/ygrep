#!/bin/bash

set -eou pipefail

if ! command -v go &> /dev/null; then
    echo "Error: Go not installed"
    exit 1
fi

go_version=$(go version)

if [[ "$go_version" =~ go([0-9]+\.[0-9]+) ]]; then
    version=${BASH_REMATCH[1]}
    major_version=$(echo $version | cut -d. -f1)
    minor_version=$(echo $version | cut -d. -f2)

    if ! (( major_version > 1 )) && { ! (( major_version == 1 )) || (( minor_version < 11 )); }; then
        echo "Error: Go version < 1.11"
        exit 1
    fi
else
    echo "Error: Go version didn't parse"
fi

go install github.com/HasanKaradmir/ygrep@latest

GOPATH=$(go env GOPATH)/bin
GOROOT=$(go env GOROOT)/bin
GOBIN=$(go env GOBIN)


goPaths=($GOPATH $GOROOT $GOBIN)

for str in ${goPaths[@]}; do
    if [ -f "$str/ygrep" ]; then
        if [ "$str" = "/usr/local/go/bin" ]; then
            break
        fi
        cp $str/ygrep /usr/local/go/bin
        chmod 755 /usr/local/go/bin/ygrep
    fi
done

echo "Success!"
