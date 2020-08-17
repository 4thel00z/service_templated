#! /bin/sh

this_dir=$(dirname "$0")
export PREFIX=/usr/local
mkdir -p $PREFIX/go/src/service_templated $PREFIX/go/src/_/builds
cp -r $this_dir/../* $PREFIX/go/src/service_templated
ln -s $PREFIX/go/src/service_templated $PREFIX/go/src/_/builds/service_templated
