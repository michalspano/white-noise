#!/usr/bin/env python3
# Generating .png from a .pgm file
import png
from sys import argv
from typing import List, final
from col import Colors


# Define main function
def main(const_factor: int = 256 // 16, _MAX: final = 3, gray_sc: bool = False) -> None:
    abort(f'Usage: {argv[0]} <input_path> <output_path>') if len(argv) != 3 else None

    INPUT, OUTPUT = argv[1], argv[2]
    abort('Invalid file extension(s)') if not check_ext(INPUT, OUTPUT) else None

    # Load data from INPUT to memory
    data: List[tuple] = load_data(INPUT, const_factor)
    dim: tuple = data[0]
    p: tuple = data[1]

    # Open output file as binary to write data and handle errors
    try:
        with open(OUTPUT, 'wb') as f:
            writer = png.Writer(dim[0], dim[1], greyscale=gray_sc)
            writer.write(f, p)
    except Exception as e:
        abort(f'{e}')


# Load data from a path
def load_data(PATH: str, factor: int) -> List[tuple]:

    """
    .pgm file format:
    P2
    <width> <height>
    <max_color_value>
    <data>
    """

    temp: list = []
    dimensions: tuple = ()
    try:
        # Load non-formatted data to memory
        temp: list = [row.strip() for row in open(PATH)]
    except FileExistsError:
        abort(f'File {PATH} does not exist.')
    try:
        # Load dimensions type int
        dimensions: tuple = tuple(map(int, temp[1].split()))
    except ValueError:
        abort('Invalid dimensions.')

    # Load img data non-formatted
    img_data: list = [idx.split() for idx in temp[3:]]
    # Parse img data to int in the correct format
    # (R, G, B) format.
    parsed_data: list = []
    for current_row in img_data:
        temp: tuple = tuple()
        for gray_shade in current_row:
            try:
                temp += tuple((int(gray_shade) * factor, ) * 3, )
            except None or ValueError:
                abort(f'Invalid data in {PATH}.')
        parsed_data.append(temp)
    return [dimensions, parsed_data]


# Log function
def abort(msg: str) -> None:
    print(f'{Colors.RED}{msg}{Colors.RESET}')
    exit(1)


# Check if INPUT has extension .pgm and OUTPUT has extension .png
def check_ext(in_p: str, out_p: str) -> bool:
    return False if not in_p.endswith('.pgm') or not out_p.endswith('.png') else True


# Invoke the main function
if __name__ == '__main__':
    main()
