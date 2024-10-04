#!/bin/bash

export GOOS="$1"
export GOARCH="$2"

VERSION="$3"
PACKAGE="$4"

PROGNAME=$(basename "$(pwd)")

mkdir -p "dist/$PROGNAME-$GOOS-$GOARCH"

go build -ldflags="-X '${PACKAGE}/cmd.Version=${VERSION}'" -o "dist/$PROGNAME-$GOOS-$GOARCH" 

cd dist

[ "$GOOS" == "windows" ] && {
    zip -qr "$PROGNAME-$VERSION-$GOOS-$GOARCH.zip" "$PROGNAME-$GOOS-$GOARCH/"
} || {
    tar zcf "$PROGNAME-$VERSION-$GOOS-$GOARCH.tar.gz" "$PROGNAME-$GOOS-$GOARCH/"
}

cd ..

rm -rf "dist/$PROGNAME-$GOOS-$GOARCH"

exit 0