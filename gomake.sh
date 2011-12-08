#!/bin/sh

if [ "$1" = "" ]
then
	exit 0
fi

gofile=$1
basename=`echo $gofile | awk -F ".go" '{ print $1 }'`
lfile="$basename.6"
6g $gofile
6l -o $basename $lfile 
