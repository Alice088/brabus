#!/bin/bash

# Не забудьте сделать скрипт исполнимым: chmod +x run_easyjson.sh

export PATH=$PATH:/home/gosha/go/bin

if [ -z "$1" ]; then
  echo "Please provide the directory path"
  exit 1
fi

DIRECTORY="$1"

if [ ! -d "$DIRECTORY" ]; then
  echo "Directory $DIRECTORY not found!"
  exit 1
fi

for file in "$DIRECTORY"/*.go; do
  base=$(basename "$file" .go)

  if [[ "$base" == *_easyjson ]]; then
    continue
  fi

  generated="$DIRECTORY/${base}_easyjson.go"
  if [ -f "$generated" ]; then
    rm -f "$generated"
    echo "Remove $generated"
  fi

  echo "Processing file: $file"
  easyjson -all "$file"
done

echo "Finish generation easyjson"
