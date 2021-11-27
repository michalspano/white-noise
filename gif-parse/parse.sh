#!/bin/sh
# Expect 2 command line arguments:
# [0] = gif-stash path
# [1] = output path

stash_path=$1
output_path=$2

# Executed from the root of the git repository
python3 ./gif-parse/bin/gif-parse.pyc $stash_path $output_path
exit 0