#!/bin/bash

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

if [ -f "$SCRIPT_DIR/venv-run.sh" ]; then
    exec "$SCRIPT_DIR/venv-run.sh" "$@"
elif [ -f "$SCRIPT_DIR/../venv-run.sh" ]; then
    exec "$SCRIPT_DIR/../venv-run.sh" "$@"
elif [ -f "$SCRIPT_DIR/../../venv-run.sh" ]; then
    exec "$SCRIPT_DIR/../../venv-run.sh" "$@"
elif [ -f "$SCRIPT_DIR/../../../venv-run.sh" ]; then
    exec "$SCRIPT_DIR/../../../venv-run.sh" "$@"
else
    echo "Error: Could not find venv-run.sh script"
    exit 1
fi
