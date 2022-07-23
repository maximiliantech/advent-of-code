from importlib.resources import read_binary
import sys

# Part 1
# read input file and return the list of measurements
def read_measurements(path):
    with open(path) as file:
        return [line.rstrip() for line in file]