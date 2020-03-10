#!/bin/bash
file="aads.txt"
lines=`cat $file`
for line in $lines; do
        ./runsync "$line"
done