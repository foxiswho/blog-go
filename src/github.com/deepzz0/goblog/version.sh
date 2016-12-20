#!/bin/sh
echo "run version.sh..."
VERSIONFILE="./version"
BRANCH=`git rev-parse --abbrev-ref HEAD`
COMMIT=`git rev-parse head`
CV=$BRANCH:${COMMIT:0:9}
printf "%s" $CV > $VERSIONFILE
cat $VERSIONFILE
