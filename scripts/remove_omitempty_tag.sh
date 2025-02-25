#!/bin/bash

CURRENT_DIR=$(pwd)

# Detect the OS
OS="$(uname)"

if [ "$OS" = "Darwin" ]; then
  # macOS
  find ${CURRENT_DIR}/genproto -name "*.pb.go" -exec sh -c 'sed -i "" "s/,omitempty//" {}' \;
else
  # Linux
  find ${CURRENT_DIR}/genproto -name "*.pb.go" -exec sh -c 'sed -i "s/,omitempty//" {}' \;
fi

echo "Removed omitempty tags from .pb.go files"
