#/bin/sh

# ANSI color codes
RED='\033[0;31m'
GREEN='\033[0;32m'
RESET='\033[0m'
ITALICS='\033[3m'
stash_path="./gif-stash/"

# Check if directory exists
if [ -d "$stash_path" ]; then
    printf "${GREEN}Removing all files from ${ITALICS}${stash_path}.\n"
    # Remove all files from the directory
    rm -rf $stash_path/*
else
    # Notify the user with an error message and create the directory
    printf "${RED}${ITALICS}${stash_path} does not exist${RESET}.\n"
    mkdir ${stash_path}
    exit 1
fi
exit 0
