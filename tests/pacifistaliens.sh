#!/bin/bash

echo "Testing for pacifist aliens"
go run main.go -aliens=$1 > test.tmp
	
if [ $(cat test.tmp |grep ^Bob|cut -f2 -d' '|uniq -c|head -n1|cut -f7 -d' ') -eq 1 ]
then 
	echo "Test successful! No pacifists! Only violent war mongers!" 
else
	echo "Test failed! Multiple aliens exist in the same city!"
fi
