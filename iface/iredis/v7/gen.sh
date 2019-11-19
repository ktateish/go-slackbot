#!/bin/sh

doit () {
	package=$1
	type=$2
	opath=$(echo $type | tr A-Z a-z).go

	interfacer -for $package.$type -as i$(basename $(dirname $package)).$type | grep -E -v '^\s+[^A-Z"]' > $opath
	chmod -x $opath
}

doit github.com/go-redis/redis/v7 Client
