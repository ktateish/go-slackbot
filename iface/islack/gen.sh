#!/bin/sh

doit () {
	package=$1
	type=$2
	opath=$(echo $type | tr A-Z a-z).go

	interfacer -for $package.$type -as i$(basename $package).$type | grep -E -v '^\s+[^A-Z"]' > $opath
	chmod -x $opath
}

doit github.com/nlopes/slack RTM
doit github.com/nlopes/slack Client
