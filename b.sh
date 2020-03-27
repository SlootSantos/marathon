#!/bin/bash

echo "Moin"

for i in {1..100}
do
   curl -N http://127.0.0.1:9999/sse &
    # echo "wa?" &
done

# while true
# do 
    sleep 100000
# done