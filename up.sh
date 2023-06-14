#!/bin/bash

cd "$(dirname "$0")"

if [ ! -f "t_app" ]; then
  go build . -o t_app
fi

./t_app
