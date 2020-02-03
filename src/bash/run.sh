#!/bin/bash
# Author: swan
# Date: 30.01.20
# Use: Type random words as fast as possible
###############################################

echo Welcome to Type-Runner! Hit enter to begin the race...; read

START=$SECONDS

for i in {1..5}; do
  WORD=$(shuf -n1 /usr/share/dict/words | awk '{print tolower($0)}')
  while [[ $INPUT != $WORD ]]; do
    echo "Word #$i - Enter: $WORD"
    read INPUT
  done
done

echo "You've done it in: $(($SECONDS - $START)) seconds"
