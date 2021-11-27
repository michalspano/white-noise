#!/usr/bin/env python3

from os import path
from sys import argv, exit
# Python3 dependencies:
import glob
import json
from PIL import Image


# Store ANSI escape sequences for colors in a class
class Colors:
    RED: str = '\033[91m'
    GREEN: str = '\033[92m'
    CYAN: str = '\033[96m'
    RESET: str = '\033[0m'


# Create the main function
def main(inc_argc: int = 2) -> None:
    # Check number of command line arguments
    abort(f'Usage: ./{argv[0]} </gif-stash/> <output_path>') if len(argv[1:]) != inc_argc else None
    input_path, output_path = argv[1], argv[2]

    # Check if the input path is a directory
    abort(f'Invalid directory `{input_path}`.') if not check_if_dir_exists(input_path) else None
    # Validate file extension

    exit(f'{Colors.RED}Invalid file extension.{Colors.RESET}') if not check_file_extension(output_path) else None

    # Read default json preferences
    gif_prefs: dict = read_json('./gif-parse/config.json')['gif-preferences']

    # Initialize GIF parsing, handle errors
    try:
        # Parse list of .png files to a gif
        img, *images = [Image.open(f) for f in sorted(glob.glob(f'{input_path}*.png'))]
        img.save(fp=output_path, format='GIF', append_images=images,
                 save_all=True, duration=int(gif_prefs['duration']),
                 loop=int(gif_prefs['loop']), version='GIF89a')
    except ValueError:
        abort('Not enough values to unpack, expected at least 1.')

    # Print success message
    print(f'[  {Colors.GREEN}OK{Colors.RESET}  ] '
          f'{Colors.CYAN}Python3: GIF parse successful: `{output_path}`{Colors.RESET}')


# Function to validate if a directory exists
def check_if_dir_exists(PATH: str) -> bool:
    return path.isdir(PATH)


# Check if path ends with a desired file extension
def check_file_extension(PATH: str, ext: str = 'gif') -> bool:
    return PATH.split('.')[-1] == ext


# Read GIF preferences from a json file
def read_json(PATH: str) -> dict:
    # Check if .json source exists
    try:
        _ = open(PATH).close()
    except FileNotFoundError:
        abort(f'{Colors.RED}JSON file `{PATH}` not found.{Colors.RESET}')

    # Return json data in a dictionary
    with open(PATH, 'r') as f:
        # Handle JSON decoding errors
        try:
            data: dict = json.load(f)
        except json.decoder.JSONDecodeError:
            abort(f'{Colors.RED}JSON file `{PATH}` is invalid.{Colors.RESET}')
    return data


# Log function
def abort(msg: str) -> None:
    print(f'{Colors.RED}{msg}{Colors.RESET}')
    exit(1)


# Invoke the main function
if __name__ == '__main__':
    main()
