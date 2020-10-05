#!/bin/bash
#set -x
START=0
END=$1 
P=$2

x=0
while [ $x -lt $END ]
do
    y=0
    while [ $y -lt $END ]
    do
        ./tour -s $END -x $x -y $y -p $P
	y=$(( $y + 1 ))
    done
    x=$(( $x + 1 ))
done
