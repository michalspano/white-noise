#!/bin/sh
# This Shell script will dump all PNGs files from the main directory
# to the subdirectory of "pngs" to be transformed to a GIF.

# ANSI color codes
RED='\033[0;31m'
GREEN='\033[0;32m'
RESET='\033[0m'
ITALICS='\033[3m'

# Function to move all `PNGs` to a recognised path
function move_pngs {
	# Get the current directory
	current_dir=$(pwd)

	# Default main directory
	main_dir="white-noise"

	# Get the last part of the current directory
	last_part=$(echo "$current_dir" | awk -F '/' '{print $NF}')

	# Define desired directory
	out="gif-stash/"

	printf "${GREEN}Checking if ${ITALICS}${out}${RESET} ${GREEN}exists...${RESET}\n"
	sleep 0.5
	# Check if `out` does not exist
	if [ ! -d "$out" ]; then
		# Create out directory
		mkdir "$out"
	fi

	# Check if the user is present in the main directory
	if [ "$last_part" == "$main_dir" ]; then
		printf "${GREEN}Moving all PNGs to ${ITALICS}${out}${RESET}\n"
		sleep 0.5
		if [ -f *.png ]; then
			mv *.png "$out"
		else
			printf "${RED}No PNG files found in the current directory.${RESET}\n"
			exit 1 
		fi
	else
		printf "${RED}Make sure that you are in the ${ITALICS}$main_dir${RESET} ${RED}directory, in order to move any PNGs.${RESET}\n"
		exit 1
	fi
	exit 0
}

move_pngs