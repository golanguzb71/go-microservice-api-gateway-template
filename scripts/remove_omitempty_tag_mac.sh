#!/bin/bash

CURRENT_DIR=$(pwd)

find "${CURRENT_DIR}/genproto" -name "*.pb.go" -exec sh -c 'sed -i "" "s/,omitempty//" "{}"' \;

echo "Removed omitempty tags from .pb.go files"