#!/usr/bin/env bash
set -euo pipefail

mkdir -p bin
rm -f bin/*

TARGET_OS=${1:-darwin}
TARGET_ARCH=${2:-arm64}

source ./install.sh $TARGET_OS $TARGET_ARCH
