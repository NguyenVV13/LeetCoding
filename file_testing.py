import sys
import os.path

# This program opens a file, then reads and prints every line in the file

# To use:
# python file_testing.py <filename>
if len(sys.argv) != 2:
    print("To use this program: `python file_testing.py <filename>`")
    exit(0)

filename = sys.argv[1]
if os.path.isfile(filename) is False:
    print("That file does not exist")
    exit(0)

with open(filename, "r") as file:
    for line in file:
        print(line.strip())