#!/bin/bash
set -xe
echo "start" >&2

#free -h >&2
#cd "$(dirname "$0")"
#ll
#if [ ! -f "t_app" ]; then
#  echo "install go" >&2
#  apk add --no-cache go=1.20-r5 >&2
#  echo "build app" >&2
#  go build . -o t_app >&2
#  echo "end build" >&2
#  free -h >&2
#fi

./t_app
echo "end" >&2
