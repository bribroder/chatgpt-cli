#!/usr/bin/env bash
set -euo pipefail

mkdir -p bin
rm -f bin/*

TARGET_OS=${1:-darwin}

source ./install.sh $TARGET_OS
