#!/bin/bash

#no forgot chmod +x run_easyjson.sh

export PATH=$PATH:/home/gosha/go/bin


if [ ! -d "../pkg/dto" ]; then
  echo "Directory ./pkg/dto not found!"
  exit 1
fi

for file in ../pkg/dto/*.go; do
  base=$(basename "$file" .go)

  if [[ "$base" == *_easyjson ]]; then
    echo "Skipping generated file: $file"
    continue
  fi

  generated="../pkg/dto/${base}_easyjson.go"
  if [ -f "$generated" ]; then
    echo "Skipping file, generated file exists: $file"
    continue
  fi

  echo "Processing file: $file"
  easyjson -all "$file"
done



echo "Finish generate easyjson"