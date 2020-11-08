#! /usr/bin/env sh

# Source:
# https://gist.githubusercontent.com/lukechilds/a83e1d7127b78fef38c2914c4ececc3c/

get_latest_release() {
  curl --silent "https://api.github.com/repos/$1/releases/latest" | # Get latest release from GitHub api
    grep '"tag_name":' |                                            # Get tag line
    sed -E 's/.*"([^"]+)".*/\1/'                                    # Pluck JSON value
}

# Usage
# $ get_latest_release "creationix/nvm"
# v0.31.4
get_latest_release $1
