#!/bin/sh

set -e

if [ ! -f "build/env.sh" ]; then
    echo "$0 must be run from the root of the repository."
    exit 2
fi

# Create fake Go workspace if it doesn't exist yet.
workspace="$PWD/build/_workspace"
root="$PWD"
xeniodir="$workspace/src/github.com/xenioplatform"
if [ ! -L "$xeniodir/go-xenio" ]; then
    mkdir -p "$xeniodir"
    cd "$xeniodir"
    ln -s ../../../../../. go-xenio
    cd "$root"
fi

# Set up the environment to use the workspace.
GOPATH="$workspace"
export GOPATH

# Run the command inside the workspace.
cd "$xeniodir/go-xenio"
PWD="$xeniodir/go-xenio"

# Launch the arguments with the configured environment.
exec "$@"
