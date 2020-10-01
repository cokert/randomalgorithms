#!/bin/zsh
START=0
END=$1 

for (( x=$START; x<$END; x++ ))
do
    for (( y=$START; y<$END; y++ ))
    do
        go run . -s $END -x $x -y $y
    done
done