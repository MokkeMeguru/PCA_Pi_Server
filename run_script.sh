#!/bin/bash
# shellcheck disable=SC1035
if !(type "wget" > /dev/null 2>&1); then
  echo "Please Install 'wget'"
  exit 1
fi

if [ ! -e "PCA_Pi_Server" ]; then
  make all
fi

export GIN_MODE=release
./PCA_Pi_Server server